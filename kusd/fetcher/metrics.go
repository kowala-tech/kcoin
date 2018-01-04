// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/kowala-tech/kUSD/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewMeter("kusd/fetcher/prop/announces/in")
	propAnnounceOutTimer  = metrics.NewTimer("kusd/fetcher/prop/announces/out")
	propAnnounceDropMeter = metrics.NewMeter("kusd/fetcher/prop/announces/drop")
	propAnnounceDOSMeter  = metrics.NewMeter("kusd/fetcher/prop/announces/dos")

	propBroadcastInMeter   = metrics.NewMeter("kusd/fetcher/prop/broadcasts/in")
	propBroadcastOutTimer  = metrics.NewTimer("kusd/fetcher/prop/broadcasts/out")
	propBroadcastDropMeter = metrics.NewMeter("kusd/fetcher/prop/broadcasts/drop")
	propBroadcastDOSMeter  = metrics.NewMeter("kusd/fetcher/prop/broadcasts/dos")

	headerFetchMeter = metrics.NewMeter("kusd/fetcher/fetch/headers")
	bodyFetchMeter   = metrics.NewMeter("kusd/fetcher/fetch/bodies")

	headerFilterInMeter  = metrics.NewMeter("kusd/fetcher/filter/headers/in")
	headerFilterOutMeter = metrics.NewMeter("kusd/fetcher/filter/headers/out")
	bodyFilterInMeter    = metrics.NewMeter("kusd/fetcher/filter/bodies/in")
	bodyFilterOutMeter   = metrics.NewMeter("kusd/fetcher/filter/bodies/out")
)
