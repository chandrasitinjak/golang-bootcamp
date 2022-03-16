package models

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	Item_Code   string `json:"itemCode" gorm:"not null;type:varchar(191)"`
	Description string `json:"description" gorm:"not null;type:varchar(191)"`
	Quantity    int    `json:"quantity"`
	OrderID     uint
}
