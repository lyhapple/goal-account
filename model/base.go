package model

import "time"

type BaseModel struct {
	ID         int       `gorm:"column:id;type:int;size:11" json:"id"`
	Deleted    bool      `gorm:"column:deleted;type:tinyint;size:1" json:"deleted"`
	Creator    int       `gorm:"column:creator;type:int;size:11" json:"creator"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}
