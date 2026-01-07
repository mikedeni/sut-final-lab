package entity

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title string  `valid:"stringlength(3|100)"`
	Price float64 `valid:"range(50|5000)"`
	Code  string  `valid:"matches(BK^[0-9]{6})"`
}
