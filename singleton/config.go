package singleton

import (
	"sync"
	"time"
)

// Config collects command args while using other patterns
type Config struct {
	Filename string
	Workers  int
	Duration time.Duration
}

var config *Config
var once sync.Once

// GetConfig from command args
func GetConfig(filename string, workers int, duration time.Duration) *Config {
	once.Do(func() {
		config = &Config{
			Filename: filename,
			Workers:  workers,
			Duration: duration,
		}
	})
	return config
}
