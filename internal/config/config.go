package config

import (
	"github.com/TypicalAM/goread/internal/backend"
	"github.com/TypicalAM/goread/internal/colorscheme"
)

// Config is the configuration for the program
type Config struct {
	Colors  colorscheme.Colorscheme
	Backend backend.Backend
}

// New returns a new Config
func New(colors colorscheme.Colorscheme, urlPath, cachePath string, resetCache bool) (Config, error) {
	// Create a new config
	config := Config{}

	// Set the colorscheme
	config.Colors = colors

	// Get the backend
	backend, err := backend.New(urlPath, cachePath, resetCache)
	if err != nil {
		return config, err
	}

	// Set the backend
	config.Backend = backend

	// Return the config
	return config, nil
}

// Close closes the config
func (c Config) Close() error {
	return c.Backend.Close()
}
