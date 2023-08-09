package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Price int64  `gorm:"type:integer" json:"price"`
}

func (u Product) TableName() string {
	return "gorm_product"
}

type ProductSerializer struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
}

func (self Product) Serializer() ProductSerializer {
	return ProductSerializer{
		ID:        self.ID,
		CreatedAt: self.CreatedAt.Truncate(time.Second),
		UpdatedAt: self.UpdatedAt.Truncate(time.Second),
		Name:      self.Name,
		Price:     self.Price,
	}
}
