package config

import (
	storage "github.com/ygpkg/storage-go"

	_ "github.com/ygpkg/storage-go/driver/cos"
	_ "github.com/ygpkg/storage-go/driver/local"
	_ "github.com/ygpkg/storage-go/driver/minio"
	_ "github.com/ygpkg/storage-go/driver/seaweedfs"
)

type Config struct {
	storage.Config `yaml:",inline"`

	Driver     string `yaml:"driver"`
	ServerPort string `yaml:"server_port"`
}

func (c *Config) Defaults() {
	if c.ServerPort == "" {
		c.ServerPort = ":8080"
	}
}

func NewStorage(cfg Config) (storage.Storage, error) {
	cfg.Defaults()
	if cfg.Driver == "" {
		return nil, storage.ErrInvalidConfig
	}
	return storage.New(cfg.Driver, cfg.Config)
}
