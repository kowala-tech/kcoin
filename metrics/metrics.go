// Package metrics provides general system and process level metrics collection.
package metrics

import (
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/kowala-tech/kcoin/log"
	"github.com/rcrowley/go-metrics/exp"

	prometheusmetrics "github.com/kowala-tech/go-metrics-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// MetricsEnabledFlag is the CLI flag name to use to enable metrics collections.
	MetricsEnabledFlag = "metrics"

	// MetricsPrometheusAddressFlag is the CLI flag name to use to set the Prometheus server address
	MetricsPrometheusAddressFlag = "metrics-prometheus-address"

	// MetricsPrometheusSubsystemFlag is the CLI flag name to use to set the Prometheus subsystem name
	MetricsPrometheusSubsystemFlag = "metrics-prometheus-subsystem"

	DashboardEnabledFlag = "dashboard"
)

// Enabled is the flag specifying if metrics are enable or not.
var Enabled = false

// Init enables or disables the metrics system. Since we need this to run before
// any other code gets to create meters and timers, we'll actually do an ugly hack
// and peek into the command line args for the metrics flag.
func init() {
	for _, arg := range os.Args {
		if flag := strings.TrimLeft(arg, "-"); flag == MetricsEnabledFlag || flag == DashboardEnabledFlag {
			log.Info("Enabling metrics collection")
			Enabled = true
		}
	}

	exp.Exp(DefaultRegistry)
}

// CollectProcessMetrics periodically collects various metrics about the running
// process.
func CollectProcessMetrics(refresh time.Duration, promAddr, promSubSys string) {
	// Short circuit if the metrics system is disabled
	if !Enabled {
		return
	}

	// Set up Prometheus
	go func() {
		prometheusRegistry := prometheus.DefaultGatherer
		metricsRegistry := DefaultRegistry
		pClient := prometheusmetrics.NewPrometheusProvider(metricsRegistry, "eth", promSubSys, (prometheusRegistry).(*prometheus.Registry), refresh)
		go pClient.UpdatePrometheusMetrics()

		log.Info("Starting Prometheus metrics", "address", promAddr, "subsystem", promSubSys)
		http.Handle("/metrics", promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{}))
		http.ListenAndServe(promAddr, nil)
	}()

	// Create the various data collectors
	memstats := make([]*runtime.MemStats, 2)
	diskstats := make([]*DiskStats, 2)
	for i := 0; i < len(memstats); i++ {
		memstats[i] = new(runtime.MemStats)
		diskstats[i] = new(DiskStats)
	}
	// Define the various metrics to collect
	memAllocs := GetOrRegisterMeter("system/memory/allocs", DefaultRegistry)
	memFrees := GetOrRegisterMeter("system/memory/frees", DefaultRegistry)
	memInuse := GetOrRegisterMeter("system/memory/inuse", DefaultRegistry)
	memPauses := GetOrRegisterMeter("system/memory/pauses", DefaultRegistry)

	var diskReads, diskReadBytes, diskWrites, diskWriteBytes Meter
	if err := ReadDiskStats(diskstats[0]); err == nil {
		diskReads = GetOrRegisterMeter("system/disk/readcount", DefaultRegistry)
		diskReadBytes = GetOrRegisterMeter("system/disk/readdata", DefaultRegistry)
		diskWrites = GetOrRegisterMeter("system/disk/writecount", DefaultRegistry)
		diskWriteBytes = GetOrRegisterMeter("system/disk/writedata", DefaultRegistry)
	} else {
		log.Debug("Failed to read disk metrics", "err", err)
	}
	// Iterate loading the different stats and updating the meters
	for i := 1; ; i++ {
		runtime.ReadMemStats(memstats[i%2])
		memAllocs.Mark(int64(memstats[i%2].Mallocs - memstats[(i-1)%2].Mallocs))
		memFrees.Mark(int64(memstats[i%2].Frees - memstats[(i-1)%2].Frees))
		memInuse.Mark(int64(memstats[i%2].Alloc - memstats[(i-1)%2].Alloc))
		memPauses.Mark(int64(memstats[i%2].PauseTotalNs - memstats[(i-1)%2].PauseTotalNs))

		if ReadDiskStats(diskstats[i%2]) == nil {
			diskReads.Mark(diskstats[i%2].ReadCount - diskstats[(i-1)%2].ReadCount)
			diskReadBytes.Mark(diskstats[i%2].ReadBytes - diskstats[(i-1)%2].ReadBytes)
			diskWrites.Mark(diskstats[i%2].WriteCount - diskstats[(i-1)%2].WriteCount)
			diskWriteBytes.Mark(diskstats[i%2].WriteBytes - diskstats[(i-1)%2].WriteBytes)
		}
		time.Sleep(refresh)
	}
}
