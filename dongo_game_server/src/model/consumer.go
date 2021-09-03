package model

// 顾客
type Consumer struct {
	Id        int    `gorm:"primary_key" json:"id"`       // 主键ID
	Name      string `gorm:"not null" json:"name"`        // 用户名称
	ProjectId int    `gorm:"default:0" json:"project_id"` // 项目id
	WxOpenId  string `gorm:"not null" json:"wx_open_id"`  // 微信openid
	Ct        int64  `gorm:"default:0" json:"ct"`         // 创建时间
	Mt        int64  `gorm:"default:0" json:"mt"`         // 修改时间
}
