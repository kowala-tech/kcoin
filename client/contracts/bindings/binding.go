package bindings

import (
	"github.com/kowala-tech/kcoin/client/common"
)

type Binding interface {
	Address() common.Address
}
