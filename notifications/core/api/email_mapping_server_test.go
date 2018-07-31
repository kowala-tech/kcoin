package api

import (
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/kowala-tech/kcoin/notifications/keyvalue"
	"github.com/stretchr/testify/require"
	"github.com/yourheropaul/inj"
	"golang.org/x/net/context"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/notifications/protocolbuffer"
	"google.golang.org/grpc"
)

func getFreePort(t *testing.T) int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	require.NoError(t, err)

	l, err := net.ListenTCP("tcp", addr)
	require.NoError(t, err)
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

func setup(t *testing.T) (keyvalue.KeyValue, protocolbuffer.EmailMappingClient, *grpc.Server, *grpc.ClientConn) {
	kv := keyvalue.NewMemoryKeyValue()
	apiServer := NewEmailMappingServer(logger)
	p := &mockedPersistance{}

	gr := inj.NewGraph()
	gr.Provide(
		apiServer,
		kv,
		p,
	)

	valid, messages := gr.Assert()
	require.True(t, valid, messages)

	port := getFreePort(t)
	addr := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", addr)
	require.NoError(t, err)

	grpcServer := grpc.NewServer()
	protocolbuffer.RegisterEmailMappingServer(grpcServer, apiServer)
	go grpcServer.Serve(lis)
	time.Sleep(10 * time.Millisecond)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)

	client := protocolbuffer.NewEmailMappingClient(conn)

	return kv, client, grpcServer, conn
}

func TestServer_RegistersStoresInKV(t *testing.T) {
	wallet := "abcde"
	email := "test@test.com"

	kv, apiClient, grpcServer, grpcClient := setup(t)
	defer grpcServer.Stop()
	defer grpcClient.Close()

	// Make sure initial data is valid
	val, err := kv.GetString(wallet)
	require.NoError(t, err)
	require.Equal(t, val, "")

	_, err = apiClient.Register(context.Background(), &protocolbuffer.RegisterRequest{Wallet: wallet, Email: email})
	require.NoError(t, err)

	// Make sure data changed
	val, err = kv.GetString(wallet)
	require.NoError(t, err)
	require.Equal(t, val, email)
}

func TestServer_RegistersStoresInKVUsingLowercases(t *testing.T) {
	wallet := "abcdeABCDE"
	email := "test@test.com"

	kv, apiClient, grpcServer, grpcClient := setup(t)
	defer grpcServer.Stop()
	defer grpcClient.Close()

	// Make sure initial data is valid
	val, err := kv.GetString(wallet)
	require.NoError(t, err)
	require.Equal(t, val, "")

	_, err = apiClient.Register(context.Background(), &protocolbuffer.RegisterRequest{Wallet: wallet, Email: email})
	require.NoError(t, err)

	// Make sure data changed
	val, err = kv.GetString(strings.ToLower(wallet))
	require.NoError(t, err)
	require.Equal(t, val, email)
}

func TestServer_RegistersFailsIfExists(t *testing.T) {
	wallet := "abcde"
	email1 := "test1@test.com"
	email2 := "test2@test.com"

	kv, apiClient, grpcServer, grpcClient := setup(t)
	defer grpcServer.Stop()
	defer grpcClient.Close()

	_, err := apiClient.Register(context.Background(), &protocolbuffer.RegisterRequest{Wallet: wallet, Email: email1})
	require.NoError(t, err)
	_, err = apiClient.Register(context.Background(), &protocolbuffer.RegisterRequest{Wallet: wallet, Email: email2})
	require.Error(t, err)

	// Make sure data changed once
	val, err := kv.GetString(wallet)
	require.NoError(t, err)
	require.Equal(t, val, email1)
}

func TestServer_UnregistersRemovesFromKV(t *testing.T) {
	wallet := "abcde"
	email := "test@test.com"

	kv, apiClient, grpcServer, grpcClient := setup(t)
	defer grpcServer.Stop()
	defer grpcClient.Close()

	_, err := apiClient.Register(context.Background(), &protocolbuffer.RegisterRequest{Wallet: wallet, Email: email})
	require.NoError(t, err)

	// Make sure initial data is valid
	val, err := kv.GetString(wallet)
	require.NoError(t, err)
	require.Equal(t, val, email)

	_, err = apiClient.Unregister(context.Background(), &protocolbuffer.UnregisterRequest{Wallet: wallet})
	require.NoError(t, err)

	// Make sure data changed
	val, err = kv.GetString(wallet)
	require.NoError(t, err)
	require.Equal(t, val, "")
}

func TestServer_UnregisterFailsIfDoesNotExist(t *testing.T) {
	wallet := "abcde"

	_, apiClient, grpcServer, grpcClient := setup(t)
	defer grpcServer.Stop()
	defer grpcClient.Close()

	_, err := apiClient.Unregister(context.Background(), &protocolbuffer.UnregisterRequest{Wallet: wallet})
	require.Error(t, err)
}

type mockedPersistance struct {
}

func (*mockedPersistance) Save(tx *protocolbuffer.Transaction) error {
	panic("implement me")
}

func (*mockedPersistance) GetTxByHash(hash common.Hash) (*protocolbuffer.Transaction, error) {
	panic("implement me")
}

func (*mockedPersistance) GetTxsFromAccount(address common.Address) ([]*protocolbuffer.Transaction, error) {
	panic("implement me")
}
