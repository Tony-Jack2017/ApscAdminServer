package model

import (
	"ApscAdmin/common/config"
	"ApscAdmin/common/model"
)

type Account struct {
	AccountID   string `gorm:"primary_key;not null" json:"account_id"` //ID
	Account     string `gorm:"" json:"account"`                        //账号
	Email       string `gorm:"" json:"email"`                          //邮箱
	Phone       string `gorm:"" json:"phone"`                          //手机
	Password    string `gorm:"" json:"password"`                       //密码
	AccountType string `gorm:"" json:"account_type"`                   //账号类型
	ShowEmail   bool   `gorm:"" json:"show_email"`                     //邮箱号展示
	ShowPhone   bool   `gorm:"" json:"show_phone"`                     //手机号展示
	model.TimeModel
}

func (account *Account) TableName() string {
	return config.Conf.System.Name + "account"
}

// API请求结构

type AccountSignUp struct {
}

type AccountSignIn struct {
	LoginType string `json:"login_type"` //	登录方式
	Account   string `json:"account"`    // 账号
	Password  string `json:"password"`   // 密码
}
