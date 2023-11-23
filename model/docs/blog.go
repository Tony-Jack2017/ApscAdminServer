package docs

import (
	"ApscAdmin/model"
	"time"
)

type Blog struct {
	BlogID        int16       `gorm:"primary_key" json:"blog_id"`       //		博客ID
	Title         string      `gorm:"varchar(255)" json:"title"`        //		博客标题
	Description   string      `gorm:"varchar(1024)" json:"description"` //	博客描述
	Cover         string      `gorm:"varchar(255)" json:"cover"`        //	博客封面
	Type          string      `gorm:"varchar(322)" json:"type"`         //	博客类型
	PublishedTime time.Time   `json:"published_time"`                   //	博客发布时间
	Author        model.Admin `json:"author"`                           //	博客作者
}
