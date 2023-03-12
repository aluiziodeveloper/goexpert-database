package main

import (
	"fmt"

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
	// Select One
	var sale Sale
	db.First(&sale, "description = ?", "Product 04")
	fmt.Println("Select One:")
	fmt.Println(sale)
	// Select All
	var allSales []Sale
	db.Find(&allSales)
	fmt.Println("Select All:")
	for _, sale := range allSales {
		fmt.Println(sale)
	}
	// Select, Limit
	db.Limit(2).Find(&allSales)
	fmt.Println("Select Limit 2:")
	for _, sale := range allSales {
		fmt.Println(sale)
	}
	// Select, Limit, Offset
	db.Limit(2).Offset(2).Find(&allSales)
	fmt.Println("Select Limit 2 Offset 2:")
	for _, sale := range allSales {
		fmt.Println(sale)
	}
	// Select, Where
	db.Where("price > ?", 500).Find(&allSales)
	fmt.Println("Select Where price > 500:")
	for _, sale := range allSales {
		fmt.Println(sale)
	}
	// Select, Where, Like
	db.Where("description LIKE ?", "%Product 03%").Find(&allSales)
	fmt.Println("Select Where description LIKE %Product 03%:")
	for _, sale := range allSales {
		fmt.Println(sale)
	}
}
