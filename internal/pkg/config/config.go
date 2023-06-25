package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/errors"
	"gopkg.in/yaml.v3"
)

var (
	cfg *model.Config
)

// SetConfig 设置配置
func SetConfig(file string) error {
	f, err := filepath.Abs(file)
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

// GetConfig 获取配置
func GetConfig() *model.Config {
	if cfg == nil {
		return &model.Config{}
	}

	return cfg
}
