// Go port of Coda Hale's Metrics library
//
// <https://github.com/rcrowley/go-metrics>
//
// Coda Hale's original work: <https://github.com/codahale/metrics>
package metrics

import (
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	prometheusmetrics "github.com/kowala-tech/go-metrics-prometheus"
	"github.com/kowala-tech/kcoin/client/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Enabled is checked by the constructor functions for all of the
// standard metrics.  If it is true, the metric returned is a stub.
//
// This global kill-switch helps quantify the observer effect and makes
// for less cluttered pprof profiles.
var Enabled bool = false

const (
	// MetricsEnabledFlag is the CLI flag name to use to enable metrics collections.
	MetricsEnabledFlag = "metrics"

	// MetricsPrometheusAddressFlag is the CLI flag name to use to set the Prometheus server address
	MetricsPrometheusAddressFlag = "metrics-prometheus-address"

	// MetricsPrometheusSubsystemFlag is the CLI flag name to use to set the Prometheus subsystem name
	MetricsPrometheusSubsystemFlag = "metrics-prometheus-subsystem"

	DashboardEnabledFlag = "dashboard"
)

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
		location1 := i % 2
		location2 := (i - 1) % 2

		runtime.ReadMemStats(memstats[location1])
		memAllocs.Mark(int64(memstats[location1].Mallocs - memstats[location2].Mallocs))
		memFrees.Mark(int64(memstats[location1].Frees - memstats[location2].Frees))
		memInuse.Mark(int64(memstats[location1].Alloc - memstats[location2].Alloc))
		memPauses.Mark(int64(memstats[location1].PauseTotalNs - memstats[location2].PauseTotalNs))

		if ReadDiskStats(diskstats[location1]) == nil {
			diskReads.Mark(diskstats[location1].ReadCount - diskstats[location2].ReadCount)
			diskReadBytes.Mark(diskstats[location1].ReadBytes - diskstats[location2].ReadBytes)
			diskWrites.Mark(diskstats[location1].WriteCount - diskstats[location2].WriteCount)
			diskWriteBytes.Mark(diskstats[location1].WriteBytes - diskstats[location2].WriteBytes)
		}
		time.Sleep(refresh)
	}
}
