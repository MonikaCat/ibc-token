package init

import (
	"github.com/spf13/cobra"

	config "github.com/MonikaCat/ibcjuno/utils"
)

// ConfigCreator represents a function that builds a Config instance from the flags that have been specified by the
// user inside the given command.
type ConfigCreator = func(cmd *cobra.Command) config.Config

// DefaultConfigCreator represents the default configuration creator that builds a Config instance using the values
// specified using the default flags.
func DefaultConfigCreator(_ *cobra.Command) config.Config {
	return config.DefaultConfig()
}

// Config contains the configuration data for the init command
type Config struct {
	createConfig ConfigCreator
}

// NewConfig allows to build a new Config instance
func NewConfig() *Config {
	return &Config{}
}

// GetConfigCreator return the function that should be run to create a configuration from a set of
// flags specified by the user with the "init" command
func (c *Config) GetConfigCreator() ConfigCreator {
	if c.createConfig == nil {
		return DefaultConfigCreator
	}
	return c.createConfig
}
