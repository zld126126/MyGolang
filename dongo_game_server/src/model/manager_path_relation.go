package model

// 管理员和路径关系
type ManagerPathRelation struct {
	Id          int `gorm:"primary_key" json:"id"`       // 主键ID
	ManagerId   int `gorm:"default:0" json:"manager_id"` // 管理员id
	ManagerPath int `gorm:"default:0" json:"manager_id"` // 路径id
}
