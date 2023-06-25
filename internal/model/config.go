package model

// Config 业务配置
type Config struct {
	Database DatabaseConfigs `yaml:"database"`
}

// DatabaseConfigs 数据库配置
type DatabaseConfigs struct {
	UserInfo DatabaseConfig `yaml:"user_info"`
}

// DatabaseConfig database 配置
type DatabaseConfig struct {
	Type DatabaseType `yaml:"type"`

	Namespace string `yaml:"namespace"`
	Target    string `yaml:"target"`
	Name      string `yaml:"name"`
}

// DatabaseType database 类型
type DatabaseType string

const (
	// DatabaseTypeMysql mysql
	DatabaseTypeMysql = "mysql"
)
