package exchange

import "github.com/kowala-tech/kcoin/mock-exchange/server"

type Transformer interface {
	Transform(request server.FetchDataRequest) string
}
