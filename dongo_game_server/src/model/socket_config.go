package model

type SocketStatus int32 // socket运行状态

const (
	SocketStatusUnknown SocketStatus = iota // 未知状态
	SocketStatusRun                         // 启动
	SocketStatusStop                        // 停止
)

// socket 配置类
type SocketConfig struct {
	Id        int64        `gorm:"primary_key" json:"id"`      // 主键ID
	Port      int64        `gorm:"default:0" json:"port"`      // 端口号 20000-30000
	ProjectId int64        `gorm:"default:0" json:"projectId"` // 项目id
	Status    SocketStatus `gorm:"default:0" json:"status"`    // socket运行状态
	Ct        int64        `gorm:"default:0" json:"ct"`        // 创建时间
	Mt        int64        `gorm:"default:0" json:"mt"`        // 修改时间
}
