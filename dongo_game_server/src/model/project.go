package model

type ProjectStatus int32

const (
	ProjectStatus_Unknown ManagerType = iota // 未知状态
	ProjectStatus_Run                        // 启动状态
	ProjectStatus_Stop                       // 运行状态
)

// 项目类
type Project struct {
	Id           int           `gorm:"primary_key" json:"id"`          // 主键ID
	Name         string        `gorm:"not null" json:"name"`           // 项目名称
	Token        string        `gorm:"not null" json:"token"`          // Token刷新令牌
	SocketPort   int           `gorm:"default:0" json:"socket_port"`   // 项目Socket端口
	RestApi      string        `gorm:"default:''" json:"rest_api"`     // HttpRestApi 例如/:id
	SocketStatus ProjectStatus `gorm:"default:0" json:"socket_status"` // socket运行状态
	ResourcePath string        `gorm:"not null" json:"resource_path"`  // 静态资源根目录
	Ct           int64         `gorm:"default:0" json:"ct"`            // 创建时间
	Dt           int64         `gorm:"default:0" json:"dt"`            // 删除时间
	Mt           int64         `gorm:"default:0" json:"mt"`            // 修改时间
}

func (p *Project) IsDt() bool {
	return p.Dt != 0
}
