package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	gorm.Model
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
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
	db.AutoMigrate(Product{})

	// Create Category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// Create Product
	db.Create(&Product{
		Name:       "Impressora",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	// Select All: com dados relacionados
	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
