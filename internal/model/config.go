package model

// Config 业务配置
type Config struct {
	Database DatabaseConfigs `yaml:"database"`

	JWT JWTConfig `yaml:"jwt"`
}

// JWTConfig jwt 相关配置
type JWTConfig struct {
	SignKey string `yaml:"sign_key"`
}

// DatabaseConfigs 数据库配置
type DatabaseConfigs struct {
	UserInfo      DatabaseConfig `yaml:"user_info"`
	Poi           DatabaseConfig `yaml:"poi"`
	WxMiniProgram DatabaseConfig `yaml:"wx_mini_program"`
}

// DatabaseConfig database 配置
type DatabaseConfig struct {
	Type DatabaseType `yaml:"type"`

	Namespace string `yaml:"namespace"`
	Target    string `yaml:"target"`
	Name      string `yaml:"name"`
}

// DatabaseType database 类型
type DatabaseType = string

const (
	// DatabaseTypeMysql mysql
	DatabaseTypeMysql = "mysql"
	// DatabaseTypeWxMiniProgram 微信小程序
	DatabaseTypeWxMiniProgram = "wx_mini_program"
	// DatabaseTypeTxMap 腾讯地图
	DatabaseTypeTxMap = "tx_map"
)
