package models

type Product struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"type:varchar(100)"`
	Price      float64 `gorm:"not null"`
	Stock      int     `gorm:"default:0"`
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
}
