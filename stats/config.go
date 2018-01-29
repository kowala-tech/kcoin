package stats

import (
	"bytes"
	"html/template"
	"os"
)

type Config struct {
	URL string `toml:",omitempty"`
}

func (config *Config) GetURL() string {
	tpl, err := template.New("url").Parse(config.URL)
	if err != nil {
		// If there's an error return the original url
		return config.URL
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, map[string]string{
		"Hostname": hostname(),
	})
	if err != nil {
		// If there's an error return the original url
		return config.URL
	}

	return buf.String()
}

func hostname() string {
	value, err := os.Hostname()
	if err != nil {
		return "unknown-hostname"
	}
	return value
}
