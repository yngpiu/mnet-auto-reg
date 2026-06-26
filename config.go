package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Password string `json:"password"`
	Mode     string `json:"mode"`
}

var configFile string

func init() {
	exe, _ := os.Executable()
	configFile = filepath.Join(filepath.Dir(exe), "config.json")
}

func defaultConfig() Config {
	return Config{
		Password: "nmixx0222-",
		Mode:     "auto",
	}
}

func loadConfig() Config {
	data, err := os.ReadFile(configFile)
	if err != nil {
		cfg := defaultConfig()
		saveConfig(cfg)
		return cfg
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		cfg := defaultConfig()
		saveConfig(cfg)
		return cfg
	}

	def := defaultConfig()
	if cfg.Password == "" {
		cfg.Password = def.Password
	}
	if cfg.Mode == "" {
		cfg.Mode = def.Mode
	}
	return cfg
}

func saveConfig(cfg Config) {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(configFile, data, 0644)
}

func updateConfig(updates map[string]string) Config {
	cfg := loadConfig()
	if v, ok := updates["password"]; ok {
		cfg.Password = v
	}
	if v, ok := updates["mode"]; ok {
		cfg.Mode = v
	}
	saveConfig(cfg)
	return cfg
}
