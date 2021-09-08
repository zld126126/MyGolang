package model

// 测试表
type Test struct {
	Id   int    `gorm:"primary_key" json:"id"` // 主键ID
	Name string `gorm:"not null" json:"name"`  // 名称
}
