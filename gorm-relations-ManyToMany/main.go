package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
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
	db.AutoMigrate(&Product{}, &Category{})

	// Create Category
	category1 := Category{Name: "Eletronicos"}
	db.Create(&category1)
	category2 := Category{Name: "Informática"}
	db.Create(&category2)

	// Create Product
	db.Create(&Product{
		Name:       "Impressora",
		Price:      1000.00,
		Categories: []Category{category1, category2},
	})

	// ManyToMany
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("  - ", product.Name)
		}
	}
}
