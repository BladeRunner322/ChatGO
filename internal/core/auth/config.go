package core_auth

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Secret string `envconfig:"SECRET" required:"true"`
	TTL    int    `envconfig:"TTL" default:"24"`
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("JWT", &cfg); err != nil {
		return Config{}, fmt.Errorf("process JWT envconfig: %w", err)
	}

	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(fmt.Errorf("load JWT config: %w", err))
	}

	return cfg
}

func (c Config) TTLDuration() time.Duration {
	return time.Duration(c.TTL) * time.Hour
}
