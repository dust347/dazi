package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/caarlos0/env/v9"
	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/errors"
	"gopkg.in/yaml.v3"
)

var (
	cfg *model.Config = &model.Config{}
)

// SetConfig4Test ...
func SetConfig4Test(c *model.Config) {
	cfg = c
}

// SetConfig 设置配置
func SetConfig(file string) error {
	f, err := filepath.Abs(file)
	log.Println(f)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return errors.WithMsg(err, "read err")
	}
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return errors.WithMsg(err, "unmarshal err")
	}

	return nil
}

// Load 加载配置
func Load() error {
	opts := env.Options{
		Prefix: "T_",
	}

	// Load env vars.
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return errors.WithMsg(err, "parse env err")
	}

	return nil
}

// GetConfig 获取配置
func GetConfig() *model.Config {
	if cfg == nil {
		return &model.Config{}
	}

	return cfg
}
