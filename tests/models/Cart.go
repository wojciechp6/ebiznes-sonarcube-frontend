package models

type Cart struct {
	ID         uint      `gorm:"primaryKey"`
	Products   []Product `gorm:"many2many:cart_products"`
	TotalPrice float64
}
