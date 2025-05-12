package config

import (
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/parsers/yaml"
)

func Load() *Config {
	var k = koanf.New(".")
	k.Load(confmap.Provider(map[string]interface{}{
		"aut.refresh_subject": RefreshTokenSubject,
		"aut.access_subject":  AccessTokenSubject,
	}, "."), nil)

	k.Load(file.Provider("config.yaml"), yaml.Parser())
	k.Load(env.Provider("GAMEAPP_", ".", func(s string) string {
		return 	}), nil)
	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}
	return &cfg
}