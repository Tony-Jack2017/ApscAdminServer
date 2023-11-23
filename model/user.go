package model

import "ApscAdmin/common/model"

type User struct {
	UserID   string `gorm:"" json:"user_id"`  //用户ID
	Username string `gorm:"" json:"username"` //用户名称
	Avatar   string `gorm:"" json:"avatar"`   //用户头像
	model.TimeModel
}
