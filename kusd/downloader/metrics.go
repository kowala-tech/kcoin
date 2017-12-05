package downloader

import (
	"github.com/kowala-tech/kUSD/metrics"
)

var (
	headerInMeter      = metrics.NewMeter("kusd/downloader/headers/in")
	headerReqTimer     = metrics.NewTimer("kusd/downloader/headers/req")
	headerDropMeter    = metrics.NewMeter("kusd/downloader/headers/drop")
	headerTimeoutMeter = metrics.NewMeter("kusd/downloader/headers/timeout")

	bodyInMeter      = metrics.NewMeter("kusd/downloader/bodies/in")
	bodyReqTimer     = metrics.NewTimer("kusd/downloader/bodies/req")
	bodyDropMeter    = metrics.NewMeter("kusd/downloader/bodies/drop")
	bodyTimeoutMeter = metrics.NewMeter("kusd/downloader/bodies/timeout")

	receiptInMeter      = metrics.NewMeter("kusd/downloader/receipts/in")
	receiptReqTimer     = metrics.NewTimer("kusd/downloader/receipts/req")
	receiptDropMeter    = metrics.NewMeter("kusd/downloader/receipts/drop")
	receiptTimeoutMeter = metrics.NewMeter("kusd/downloader/receipts/timeout")

	stateInMeter   = metrics.NewMeter("kusd/downloader/states/in")
	stateDropMeter = metrics.NewMeter("kusd/downloader/states/drop")
)
