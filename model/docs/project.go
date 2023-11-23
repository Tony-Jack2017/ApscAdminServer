package docs

type Project struct {
	ProjectID   int16  `gorm:"primary_key;not null;" json:"project_id"` //项目ID
	ProjectName string `gorm:"varchar(255)" json:"project_name"`        //项目名
	Cover       string `gorm:"varchar(255)" json:"cover"`               //项目封面
	Description string `gorm:"varchar(1024)" json:"description"`        //项目描述
	Price       int    `gorm:"int" json:"price"`                        //项目价格
	Version     string `gorm:"varchar(32)" json:"version"`              //项目版本
	Type        string `gorm:"varchar(255)" json:"type"`                //项目类型
	SaleNumber  int    `gorm:"int" json:"sale_number"`                  //项目售卖量
}
