package model

// 埋点
type Track struct {
	Id             int64      `gorm:"primary_key" json:"id"`        // 主键ID
	ProjectId      int64      `gorm:"default:0" json:"project_id"`  // 项目Id
	ConsumerId     int64      `gorm:"default:0" json:"consumer_id"` // 顾客Id
	ConsumerItemId int64      `gorm:"default:0" json:"consumer_id"` // 顾客ItemId
	SourceTp       SourceType `gorm:"default:0" json:"source_tp"`   // 来源类型
	Ct             int64      `gorm:"default:0" json:"ct"`          // 创建时间

	Items    []*TrackItem `gorm:"-" json:"items"`    // 扩展字段 埋点信息类型
	Messages []string     `gorm:"-" json:"messages"` // 打点信息 1个或多个
}

// 埋点多条信息
type TrackItem struct {
	Id             int64      `gorm:"primary_key" json:"id"`          // 主键ID
	TrackId        int64      `gorm:"default:0" json:"track_id"`      // 埋点主键ID
	ProjectId      int64      `gorm:"default:0" json:"project_id"`    // 项目id
	ConsumerId     int64      `gorm:"default:0" json:"consumer_id"`   // 顾客Id
	ConsumerItemId int64      `gorm:"default:0" json:"consumer_id"`   // 顾客ItemId
	SourceTp       SourceType `gorm:"default:0" json:"source_tp"`     // 来源类型
	Ct             int64      `gorm:"default:0" json:"ct"`            // 创建时间
	MessageIndex   int64      `gorm:"default:0" json:"message_index"` // 打点消息顺序 从0开始
	Message        string     `gorm:"default:0" json:"message"`       // 打点信息
}
