package config

import (
	"strings"
)

type Config struct {
	IgnoreAbbrs []string
	Abbrs       []string
}

// Violates checks if the passed string value exists in the forbidden list of abbreviations
// If the string exists in the list, and is not lowercase or uppercase, throws false
func (c Config) Violates(match string) bool {
	for _, abbr := range c.Abbrs {
		if strings.EqualFold(abbr, string(match)) {
			if strings.ToLower(string(match)) == string(match) ||
				strings.ToUpper(string(match)) == string(match) {
				return false
			}
			return true
		}
	}
	return false
}

type Option func(config *Config)

func DefaultConfig() *Config {
	return &Config{
		IgnoreAbbrs: []string{},
		Abbrs:       []string{"id", "http", "vat"},
	}
}

func WithOptions(options ...Option) *Config {
	c := DefaultConfig()

	for _, opt := range options {
		opt(c)
	}

	return c
}

func WithAbbrs(abbrs string) Option {
	return func(config *Config) {
		config.Abbrs = make([]string, 0)
		for _, abbr := range strings.Split(abbrs, ",") {
			if abbr != "" {
				config.Abbrs = append(config.Abbrs, abbr)
			}
		}
	}
}

func WithIgnoreAbbrs(abbrs string) Option {
	return func(config *Config) {
		config.IgnoreAbbrs = make([]string, 0)
		for _, abbr := range strings.Split(abbrs, ",") {
			if abbr != "" {
				config.IgnoreAbbrs = append(config.IgnoreAbbrs, abbr)
			}
		}
	}
}
