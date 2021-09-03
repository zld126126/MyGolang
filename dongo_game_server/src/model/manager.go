package model

type ManagerType int32

const (
	ManagerType_Unknown ManagerType = iota // 未知类型
	ManagerType_Normal                     // 普通用户
	ManagerType_Super                      // 超级用户
)

// 管理员
type Manager struct {
	Id       int         `gorm:"primary_key" json:"id"`    // 主键ID
	Name     string      `gorm:"not null" json:"name"`     // 用户名称
	Password string      `gorm:"not null" json:"password"` // 用户密码
	Tp       ManagerType `gorm:"default:0" json:"tp"`      // 用户类型
	Ct       int64       `gorm:"default:0" json:"ct"`      // 创建时间
	Dt       int64       `gorm:"default:0" json:"dt"`      // 删除时间
	Mt       int64       `gorm:"default:0" json:"mt"`      // 修改时间
}

func (p *Manager) IsDt() bool {
	return p.Dt != 0
}
