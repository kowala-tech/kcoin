package stats

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigGetURLPlain(t *testing.T) {
	config := &Config{
		URL: "FooBar",
	}
	require.Equal(t, "FooBar", config.GetURL())
}

func TestConfigGetURLTemplate(t *testing.T) {
	config := &Config{
		URL: "FooBar-{{.Hostname}}",
	}
	hn, err := os.Hostname()
	require.NoError(t, err)
	require.Equal(t, "FooBar-"+hn, config.GetURL())
}
