package model

// 项目类
type Project struct {
	Id           int64  `gorm:"primary_key" json:"id"`           // 主键ID
	Name         string `gorm:"not null" json:"name"`            // 项目名称
	Token        string `gorm:"not null" json:"token"`           // Token刷新令牌
	RestApi      string `gorm:"default:''" json:"rest_api"`      // HttpRestApi 例如/:id
	ResourcePath string `gorm:"default:''" json:"resource_path"` // 静态资源根目录
	Port         int64  `gorm:"not null" json:"port"`            // 端口号
	Ct           int64  `gorm:"default:0" json:"ct"`             // 创建时间
	Dt           int64  `gorm:"default:0" json:"dt"`             // 删除时间
	Mt           int64  `gorm:"default:0" json:"mt"`             // 修改时间
}

func (p *Project) IsDt() bool {
	return p.Dt != 0
}
