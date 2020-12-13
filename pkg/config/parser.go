package config

import (
	"gopkg.in/yaml.v2"
)

// Config is the top-level struct for config files
type Config struct {
	LocalProxies map[string]LocalProxy `yaml:"localProxies"`
}

// LocalProxy is a proxy that is running on the local machine
type LocalProxy struct {
	ProxyAddr      string `yaml:"proxyAddress"`
	ConnectingAddr string `yaml:"connectingAddress"`
}

// Parse will parse from the input
func Parse(input []byte) (Config, error) {
	c := &Config{}
	err := yaml.Unmarshal(input, c)
	if err != nil {
		return *c, err
	}

	return *c, nil
}
