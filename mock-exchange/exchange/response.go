package exchange

import "github.com/kowala-tech/kcoin/mock-exchange/app"

type Transformer interface {
	Transform(request app.Request) string
}
