package models

type Product struct {
	ID    int32   `gorm:"primaryKey"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
