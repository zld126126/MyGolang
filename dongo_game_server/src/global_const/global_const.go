package global_const

import (
	"github.com/ahmetb/go-linq/v3"
)

const (
	ConfigFileAddressRelease = `../resources/config.toml` // go build正式环境用
	ConfigFileAddressDebug   = `resources/config.toml`    // goland本地启动用

	ConfigFileKey    = `configFile`
	ConfigKey        = `config`
	ConfigVersionKey = `config_version` // 本地记录版本

	FileMaxBytes = 1024 * 1024 * 2 // 最大文件容量

	ManagerKey             = `ManagerKey_%d`          // 用户session key
	ManagerLoginKey        = `ManagerLoginKey_%d`     // 用户登陆session key
	ManagerLoginSplitKey   = `||`                     // 用户登录分割
	ManagerWebHeaderKey    = `ManagerWebHeaderKey`    // web登陆 自定义HeaderKey
	ManagerClientHeaderKey = `ManagerClientHeaderKey` // client登陆 自定义HeaderKey

	ManagerPathKey = `ManagerPathKey_%d` // 用户session key

	ProjectTokenSalt = `ProjectTokenSalt` // 项目token Salt
	ProjectKey       = `ProjectKey_%d`    // 用户session key

	FakeIdDebugKey = `Fake-Id`  // fake登陆 自定义HeaderKey
	FakeIdAdmin    = `YWRtaW4=` // fake登陆 debug模式

	SocketPortMin int64 = 12021 // socket 连接最小端口号
	SocketPortMax int64 = 12030 // socket 连接最大端口号
)

var Paths = []string{"/manager", "/manager_path"}

func IsNormalPath(p string) bool {
	item := linq.From(Paths).WhereT(func(i string) bool {
		return i == p
	}).First().(string)

	return item != ""
}
