package model

type SourceType int32 // 埋点来源类型

const (
	SourceTypeUnknown       SourceType = iota // 未知
	SourceTypeWxMiniProgram                   // 微信小程序
	SourceTypeAndriod                         // 安卓
	SourceTypeIOS                             // IOS
	SourceTypeWEB                             // WEB
	SourceTypeOther                           // 其他
)

// 顾客
type Consumer struct {
	Id    int64           `gorm:"primary_key" json:"id"` // 主键ID
	Name  string          `gorm:"not null" json:"name"`  // 用户名称
	Ct    int64           `gorm:"default:0" json:"ct"`   // 创建时间
	Mt    int64           `gorm:"default:0" json:"mt"`   // 修改时间
	Items []*ConsumerItem `gorm:"-" json:"items"`        // TODO:后期 微信/安卓/IOS/WEB 多端融合记录
}

// 顾客多端信息
type ConsumerItem struct {
	Id         int64      `gorm:"primary_key" json:"id"`        // 主键ID
	ConsumerId int64      `gorm:"default:0" json:"consumer_id"` // 顾客id
	Tp         SourceType `gorm:"not null" json:"tp"`           // 顾客类型
	OpenId     string     `gorm:"default:''" json:"open_id"`    // 可以为空 例如:微信->WxOpenId
	Ct         int64      `gorm:"default:0" json:"ct"`          // 创建时间
	Mt         int64      `gorm:"default:0" json:"mt"`          // 修改时间
}
