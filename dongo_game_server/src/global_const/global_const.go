package global_const

const (
	ConfigFileAddressRelease = `../resources/config.toml` // go build正式环境用
	ConfigFileAddressDebug   = `resources/config.toml`    // goland本地启动用

	ConfigFileKey    = `configFile`
	ConfigKey        = `config`
	ConfigVersionKey = `config_version` // 本地记录版本

	FileMaxBytes = 1024 * 1024 * 2 // 最大文件容量
)