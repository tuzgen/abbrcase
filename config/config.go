package config

import (
	"fmt"
	"regexp"
	"strings"
)

type Config struct {
	IgnoredFiles []*regexp.Regexp
	Abbrs        []string
}

func (c Config) IsFileIgnored(fileName string) bool {
	fmt.Println(c.IgnoredFiles, fileName)
	for _, ignoredFile := range c.IgnoredFiles {
		if ignoredFile.Match([]byte(fileName)) {
			return true
		}
	}
	return false
}

type Option func(config *Config)

func DefaultConfig() *Config {
	return &Config{
		IgnoredFiles: []*regexp.Regexp{},
		Abbrs:        []string{"id", "http", "vat"},
	}
}

func WithOptions(options ...Option) *Config {
	c := DefaultConfig()

	for _, opt := range options {
		opt(c)
	}

	return c
}

func WithIgnoredFiles(excludes string) Option {
	return func(config *Config) {
		for _, exclude := range strings.Split(excludes, ",") {
			if exclude != "" {
				config.IgnoredFiles = append(config.IgnoredFiles, regexp.MustCompile(exclude))
			}
		}
	}
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
