// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/kowala-tech/kUSD/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewMeter("eth/fetcher/prop/announces/in")
	propAnnounceOutTimer  = metrics.NewTimer("eth/fetcher/prop/announces/out")
	propAnnounceDropMeter = metrics.NewMeter("eth/fetcher/prop/announces/drop")
	propAnnounceDOSMeter  = metrics.NewMeter("eth/fetcher/prop/announces/dos")

	propBroadcastInMeter   = metrics.NewMeter("eth/fetcher/prop/broadcasts/in")
	propBroadcastOutTimer  = metrics.NewTimer("eth/fetcher/prop/broadcasts/out")
	propBroadcastDropMeter = metrics.NewMeter("eth/fetcher/prop/broadcasts/drop")
	propBroadcastDOSMeter  = metrics.NewMeter("eth/fetcher/prop/broadcasts/dos")

	headerFetchMeter = metrics.NewMeter("eth/fetcher/fetch/headers")
	bodyFetchMeter   = metrics.NewMeter("eth/fetcher/fetch/bodies")

	headerFilterInMeter  = metrics.NewMeter("eth/fetcher/filter/headers/in")
	headerFilterOutMeter = metrics.NewMeter("eth/fetcher/filter/headers/out")
	bodyFilterInMeter    = metrics.NewMeter("eth/fetcher/filter/bodies/in")
	bodyFilterOutMeter   = metrics.NewMeter("eth/fetcher/filter/bodies/out")
)
