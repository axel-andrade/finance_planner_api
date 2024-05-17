package models

type Category struct {
	Base
	Type string `gorm:"type:varchar(20);not null;index" json:"type"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
