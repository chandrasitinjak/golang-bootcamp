package _struct

import "github.com/jinzhu/gorm"

type Person1 struct {
	gorm.Model
	First_Name string
	Last_Name  string
}
