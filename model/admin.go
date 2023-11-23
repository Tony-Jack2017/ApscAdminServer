package model

import "ApscAdmin/common/model"

type Admin struct {
	AdminID     string `gorm:"primary_key;not null;" json:"admin_id"` //管理员ID
	Username    string `gorm:"varchar(123)" json:"username"`          //管理员名称
	Avatar      string `gorm:"varchar(255)" json:"avatar"`            //管理员头像
	Location    string `gorm:"varchar(255)" json:"location"`          //管理员地址
	Completed   int    `gorm:"int" json:"completed"`                  //资料完成度
	ConnectType string `gorm:"varchar(123)" json:"connect_type"`      //联系方式
	model.TimeModel
}
