package p2p

import (
	"context"
	"errors"
	"fmt"
	"sync"

	dstore "github.com/ipfs/go-datastore"
	ipfssync "github.com/ipfs/go-datastore/sync"
	"github.com/kowala-tech/kcoin/client/log"
	pubsub "github.com/libp2p/go-floodsub"
	libp2p "github.com/libp2p/go-libp2p"
	libp2p_host "github.com/libp2p/go-libp2p-host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
)

const (
	packageName = "p2p/host"
)

var (
	errNoPrivateKey = errors.New("Host.PrivateKey must be set to a non-nil key")
)

type host struct {
	Config
	libp2p_host.Host
	*pubsub.PubSub

	log log.Logger
}

// NewHost returns a new p2p host
func NewHost(cfg Config) *host {
	return &host{Config: cfg}
}

func (h *host) Start() error {
	h.log = h.Config.Logger
	if h.log == nil {
		h.log = log.New("package", packageName)
	}

	if h.PrivateKey == nil {
		return errNoPrivateKey
	}

	ctx := context.Background()
	host, err := libp2p.New(
		ctx,
		libp2p.Identity(*h.PrivateKey),
		libp2p.ListenAddrStrings(h.ListenAddr),
		libp2p.NATPortMap(),
	)
	if err != nil {
		return err
	}
	h.Host = host

	pubsub, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		return err
	}
	h.PubSub = pubsub

	dht := dht.NewDHT(ctx, host, ipfssync.MutexWrap(dstore.NewMapDatastore()))

	if len(h.BootstrapNodes) > 0 {
		routedHost := rhost.Wrap(host, dht)
		if err := bootstrapConnect(ctx, routedHost, h.BootstrapNodes, h.Logger); err != nil {
			h.log.Error("Could not connect to the bootstrap nodes", "err", err)
		}
	}

	if h.IsBootstrapNode {
		if err := dht.Bootstrap(ctx); err != nil {
			return err
		}
	}

	h.log.Info("Listening...", "ID", host.ID().Pretty(), "addr", h.ListenAddr)

	return nil
}

func (h *host) Stop() error {
	return h.Host.Close()
}

func bootstrapConnect(ctx context.Context, ph libp2p_host.Host, peers []pstore.PeerInfo, log log.Logger) error {
	if len(peers) == 0 {
		return errors.New("not enough bootstrap peers")
	}

	log.Info("Connecting to bootstrap nodes ...", "peers", peers)

	errs := make(chan error, len(peers))
	var wg sync.WaitGroup
	for _, p := range peers {

		// performed asynchronously because when performed synchronously, if
		// one `Connect` call hangs, subsequent calls are more likely to
		// fail/abort due to an expiring context.
		// Also, performed asynchronously for dial speed.

		wg.Add(1)
		go func(p pstore.PeerInfo) {
			defer wg.Done()
			defer log.Debug("bootstrapDial", "from", ph.ID(), "bootstrapping to", p.ID)
			log.Debug("bootstrapDial", "from", ph.ID(), "bootstrapping to", p.ID)

			ph.Peerstore().AddAddrs(p.ID, p.Addrs, pstore.PermanentAddrTTL)
			if err := ph.Connect(ctx, p); err != nil {
				log.Error("Failed to bootstrap with", "ID", p.ID, "err", err)
				errs <- err
				return
			}
			log.Debug("bootstrapDialSuccess", "ID", p.ID)
			log.Info("Bootstrapped with", "ID", p.ID)
		}(p)
	}
	wg.Wait()

	// our failure condition is when no connection attempt succeeded.
	// So drain the errs channel, counting the results.
	close(errs)
	count := 0
	var err error
	for err = range errs {
		if err != nil {
			count++
		}
	}
	if count == len(peers) {
		return fmt.Errorf("Failed to bootstrap. %s", err)
	}
	return nil
}
