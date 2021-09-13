package model

// 用户表
type User struct {
	Id       int64  `gorm:"primary_key" json:"id"`    // 主键ID
	Name     string `gorm:"not null" json:"name"`     // 用户名称
	Password string `gorm:"not null" json:"password"` // 用户密码
}
