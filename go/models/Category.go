package models

type Category struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}
