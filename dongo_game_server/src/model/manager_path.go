package model

// 管理员路径 超级管理员-默认全显示
type ManagerPath struct {
	Id         int    `gorm:"primary_key" json:"id"`       // 主键ID
	Name       string `gorm:"not null" json:"name"`        // 管理员路径名称
	OptionPath string `gorm:"not null" json:"option_path"` // 操作路径
}
