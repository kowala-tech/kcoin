package params

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeCanGetFullDomain(t *testing.T) {
	domain := KNSDomain{
		node: "multisig",
		tld:  "kowala",
	}

	assert.Equal(t, "multisig.kowala", domain.FullDomain())
}
