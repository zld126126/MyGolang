package model

// 顾客
type ProjectConsumerRelation struct {
	Id         int64 `gorm:"primary_key" json:"id"`        // 主键ID
	ProjectId  int64 `gorm:"default:0" json:"project_id"`  // 项目ID
	ConsumerId int64 `gorm:"default:0" json:"consumer_id"` // 用户顾客ID
	Ct         int64 `gorm:"default:0" json:"ct"`          // 创建时间
}
