package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Sale struct {
	ID          int `gorm:"primaryKey"`
	Description string
	Price       float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(Sale{})
	// Create
	db.Create(&Sale{
		Description: "Product 01",
		Price:       1200,
	})
	// Create Many
	sales := []Sale{
		{Description: "Product 02", Price: 29.90},
		{Description: "Product 03", Price: 129.90},
		{Description: "Product 04", Price: 229.90},
	}
	db.Create(&sales)
}
