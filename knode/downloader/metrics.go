// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/kowala-tech/kcoin/metrics"
)

var (
	headerInMeter      = metrics.NewMeter("eth/downloader/headers/in")
	headerReqTimer     = metrics.NewTimer("eth/downloader/headers/req")
	headerDropMeter    = metrics.NewMeter("eth/downloader/headers/drop")
	headerTimeoutMeter = metrics.NewMeter("eth/downloader/headers/timeout")

	bodyInMeter      = metrics.NewMeter("eth/downloader/bodies/in")
	bodyReqTimer     = metrics.NewTimer("eth/downloader/bodies/req")
	bodyDropMeter    = metrics.NewMeter("eth/downloader/bodies/drop")
	bodyTimeoutMeter = metrics.NewMeter("eth/downloader/bodies/timeout")

	receiptInMeter      = metrics.NewMeter("eth/downloader/receipts/in")
	receiptReqTimer     = metrics.NewTimer("eth/downloader/receipts/req")
	receiptDropMeter    = metrics.NewMeter("eth/downloader/receipts/drop")
	receiptTimeoutMeter = metrics.NewMeter("eth/downloader/receipts/timeout")

	stateInMeter   = metrics.NewMeter("eth/downloader/states/in")
	stateDropMeter = metrics.NewMeter("eth/downloader/states/drop")
)
