package server

type Website struct {
	WebsiteID int    `gorm:"primary_key;not null;" json:"website_id"` //网站ID
	Name      string `gorm:"" json:"name"`                            //网站名字
	Domain    string `gorm:"" json:"domain"`                          //网站域名
	Logo      string `gorm:"" json:"logo"`                            //网站logo
}
