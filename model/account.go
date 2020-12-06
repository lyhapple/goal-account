package model

type Account struct {
	BaseModel
	Username  string `gorm:"column:username;type:varchar(200);size:200" json:"username"`
	Cellphone string `gorm:"column:cellphone;type:varchar(50);size:50" json:"cellphone"`
	Email     string `gorm:"column:email;type:varchar(200);size:200" json:"email"`
	Password  string `gorm:"column:password;type:varchar(200);size:200" json:"password"`
	Salt      string `gorm:"column:salt;type:varchar(20);size:20" json:"salt"`
}
