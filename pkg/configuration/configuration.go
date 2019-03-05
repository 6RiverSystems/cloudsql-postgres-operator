/*
Copyright 2019 The cloudsql-operator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package configuration

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

// Configuration is the root configuration object.
type Configuration struct {
	// Logging holds logging-related configuration options.
	Logging Logging `toml:"logging"`
}

// SetDefaults sets default values where necessary.
func (c *Configuration) SetDefaults() {
	c.Logging.SetDefaults()
}

// Logging holds logging-related configuration options.
type Logging struct {
	// Level holds the log level to use.
	Level string `toml:"level"`
}

// SetDefaults sets default values where necessary.
func (l *Logging) SetDefaults() {
	if l.Level == "" {
		l.Level = log.InfoLevel.String()
	}
}

// NewDefaultConfiguration returns a new Configuration object with default values.
func NewDefaultConfiguration() Configuration {
	c := Configuration{}
	c.SetDefaults()
	return c
}

// MustNewConfigurationFromFile attempts to parse the specified configuration file, exiting the application if it cannot be parsed.
func MustNewConfigurationFromFile(path string) Configuration {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read the configuration file: %v", err)
	}
	var r Configuration
	if err := toml.Unmarshal(b, &r); err != nil {
		log.Fatalf("failed to read the configuration file: %v", err)
	}
	r.SetDefaults()
	return r
}
