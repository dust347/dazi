package model

// Config 业务配置
type Config struct {
	Database DatabaseConfigs `yaml:"database" envPrefix:"DATABASE_"`

	JWT JWTConfig `yaml:"jwt" envPrefix:"JWT_"`
	IM  IMConfig  `yaml:"im" envPrefix:"IM_"`
}

// JWTConfig jwt 相关配置
type JWTConfig struct {
	SignKey string `yaml:"sign_key" env:"SIGN_KEY"`
}

// IMConfig im 相关配置
type IMConfig struct {
	AppID     int    `yaml:"app_id" env:"APP_ID"`
	SecretKey string `yaml:"secret_key" env:"SECRET_KEY"`
}

// DatabaseConfigs 数据库配置
type DatabaseConfigs struct {
	UserInfo      DatabaseConfig `yaml:"user_info" envPrefix:"USER_INFO_"`
	Poi           DatabaseConfig `yaml:"poi" envPrefix:"POI_"`
	WxMiniProgram DatabaseConfig `yaml:"wx_mini_program" envPrefix:"WX_"`
	Image         DatabaseConfig `yaml:"image" envPrefix:"IMAGE_"`
	IM            DatabaseConfig `yaml:"im" envPrefix:"IM_"`
}

// DatabaseConfig database 配置
type DatabaseConfig struct {
	Type DatabaseType `yaml:"type" env:"TYPE"`

	Namespace string `yaml:"namespace" env:"NAMESPACE"`
	Target    string `yaml:"target" env:"TARGET"`
	Name      string `yaml:"name" env:"NAME"`
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
	// DatabaseTypeCos cos
	DatabaseTypeCos = "cos"
	// DatabaseTypeIM im
	DatabaseTypeIM = "im"
)
