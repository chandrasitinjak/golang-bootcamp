package models

import "time"

type Order struct {
	ID            uint      `gorm:"primaryKey"`
	Customer_Name string    `json:"customer_name" gorm:"not null;type:varchar(191)"`
	Item          []Item    `json:"items" gorm:"foreignKey:OrderID"`
	OrderedAt     time.Time `json:"createAt" gorm:"type:datetime"`
}

//type CreateOrder struct {
//	Customer_Name string `json:"customer_name" binding:"required"`
//	Item          []Item `json:"items" binding:"required"`
//}

type CreateOrder struct {
	Customer_Name string    `json:"customer_name" binding:"required"`
	OrderedAt     time.Time `json:"createAt" binding:"required"`
	Item          []Item    `json:"items" gorm:"-"`
}
