package cfg

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config holds the available configurable options
type Config struct {
	FPS float64
}

// GetConfig parses the config.tmol file
func GetConfig() Config {
	var conf Config
	if _, err := toml.DecodeFile("conf.toml", &conf); err != nil {
		var f *os.File
		if f, err = os.Create("conf.toml"); err != nil {
			log.Fatalln(err)
		}
		setDefault(&conf)
		enc := toml.NewEncoder(f)
		enc.Encode(&conf)
	}
	return conf
}

// Config defaults go here
func setDefault(c *Config) {
	c.FPS = 60.0
}
