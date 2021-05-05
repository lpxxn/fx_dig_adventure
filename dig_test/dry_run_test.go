package dig_test

import (
	"log"
	"os"
	"testing"

	"go.uber.org/dig"
)

func TestDryRun1(t *testing.T) {
	// Dry Run
	c := dig.New(dig.DryRun(true))

	type Config struct {
		Prefix string
	}
	err := c.Provide(func() (*Config, error) {
		return &Config{Prefix: "[foo] "}, nil
	})
	if err != nil {
		panic(err)
	}
	err = c.Provide(func(cfg *Config) *log.Logger {
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	if err != nil {
		panic(err)
	}
	err = c.Invoke(func(l *log.Logger) {
		l.Print("You've been invoked")
	})
	if err != nil {
		panic(err)
	}
}