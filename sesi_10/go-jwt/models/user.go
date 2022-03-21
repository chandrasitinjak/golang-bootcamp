package models

import (
	"github.com/asaskevich/govalidator"
	"go-jwt/helpers"
	"gorm.io/gorm"
	"log"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Full name is required"`
	Email    string    `gorm:"not null" json:"email" form:"email" valid:"required~Email is required, email"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Password is required, minstringlength(6)~Length minimum password is 6"`
	Products []Product `gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		log.Println(err)
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
