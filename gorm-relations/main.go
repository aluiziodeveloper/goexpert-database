package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
	gorm.Model
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
	gorm.Model
}

const dsn string = "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Create Category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// Create Product
	db.Create(&Product{
		Name:       "Impressora",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	// Create SerialNumber
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	// Select All: com dados relacionados
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, category.Name, product.SerialNumber.Number)
	}

	// HasMany
	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("  - ", product.Name, " - SerialNumber: ", product.SerialNumber.Number)
		}
	}
}
