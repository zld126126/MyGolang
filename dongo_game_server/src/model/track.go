package model

// 埋点
type Track struct {
	Id         int64  `gorm:"primary_key" json:"id"`        // 主键ID
	Name       string `gorm:"not null" json:"name"`         // 用户名称
	ProjectId  int64  `gorm:"default:0" json:"project_id"`  // 项目id
	ConsumerId int64  `gorm:"default:0" json:"consumer_id"` // 顾客id
	Ct         int64  `gorm:"default:0" json:"ct"`          // 创建时间
}
