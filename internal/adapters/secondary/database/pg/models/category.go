package models

type Category struct {
	Base
	Name string `gorm:"type:varchar(255)" json:"name"`
}
