// Package cfg implements configuration file encoding, decoding, and storing.
package cfg

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config holds the available configurable options.
type Config struct {
	FPS float64 `toml:"fps"`
}

// GetConfig parses the config.tmol file.
func GetConfig() Config {
	var conf Config
	if _, err := toml.DecodeFile("conf.toml", &conf); err != nil {
		var f *os.File
		if f, err = os.Create("conf.toml"); err != nil {
			log.Fatalln(err)
		}
		setDefault(&conf)
		enc := toml.NewEncoder(f)
		if err = enc.Encode(&conf); err != nil {
			log.Fatalln(err)
		}
	}
	return conf
}

// setDefault sets defaults for the given Config.
func setDefault(c *Config) {
	c.FPS = 60.0
}
