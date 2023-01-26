package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
)

type Config struct {
	AppPort               int      `koanf:"APP_PORT"`
	AppHost               string   `koanf:"APP_HOST" `
	SublimeSigningKey     string   `koanf:"SUBLIME_SIGNING_KEY"`
	SublimeHMACExpiration int      `koanf:"SUBLIME_HMAC_EXPIRATION"`
	SublimeURL            string   `koanf:"SUBLIME_URL"`
	SublimeApiKey         string   `koanf:"SUBLIME_API_KEY"`
	SublimeApiURL         string   `koanf:"SUBLIME_API_URL"`
	Debug                 bool     `koanf:"DEBUG"`
	TheHiveURL            string   `koanf:"THEHIVE_URL"`
	THeHiveKey            string   `koanf:"THEHIVE_KEY"`
	TheHiveAlertType      string   `koanf:"THEHIVE_ALERT_TYPE"`
	TheHiveAlertTags      []string `koanf:"THEHIVE_ALERT_TAGS"`
}

func New() (conf Config, err error) {
	newConfig := &Config{}
	if err := newConfig.Setup(); err != nil {
		return *newConfig, err
	}
	return *newConfig, nil
}

func (d *Config) Setup() error {
	k := koanf.New(".")

	//Default Values
	k.Load(confmap.Provider(map[string]interface{}{
		"THEHIVE_URL":             "http://localhost:9000",
		"THEHIVE_KEY":             "NO_KEY",
		"DEBUG":                   false,
		"SUBLIME_HMAC_EXPIRATION": 3,
		"APP_PORT":                4000,
		"APP_HOST":                "0.0.0.0",
		"SUBLIME_URL":             "http://localhost:3000",
		"THEHIVE_ALERT_TYPE":      "Phishing",
	}, "."), nil)

	k.Load(env.ProviderWithValue("", ".", func(s string, v string) (string, interface{}) {
		if strings.Contains(v, ",") {
			return s, strings.Split(v, ",")
		}
		return s, v
	}), nil)

	k.Unmarshal("", &d)
	return nil
}

func (d *Config) Print() string {
	if d != nil {
		return fmt.Sprintf("%+v", d)
	}
	return ""
}
