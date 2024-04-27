package models

type Category struct {
	Base
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
