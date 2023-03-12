package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Sale struct {
	ID          int `gorm:"primaryKey"`
	Description string
	Price       float64
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
	db.AutoMigrate(Sale{})

	// Create
	create(&Sale{
		Description: "Product 01",
		Price:       1200,
	})

	// Update
	var s Sale
	ID := 1
	updateDescription(&s, ID, "New description")

	// Delete
	var s2 Sale
	remove(&s2, ID)

	// // Create Many
	// sales := []Sale{
	// 	{Description: "Notebook dell", Price: 3929.90},
	// 	{Description: "MacBook Pro", Price: 9129.90},
	// 	{Description: "Monitor", Price: 829.90},
	// 	{Description: "Keyboard", Price: 29.90},
	// 	{Description: "Mouse", Price: 29.90},
	// }
	// createMany(&sales)

	// // Select One
	// var sale Sale
	// findByDescription(&sale, "Monitor")
	// fmt.Println("Select One:")
	// fmt.Println(sale)

	// // Select All
	// var allSales []Sale
	// findAll(&allSales)
	// fmt.Println("Select All:")
	// for _, sale := range allSales {
	// 	fmt.Println(sale)
	// }

	// // Select, Limit, Offset
	// paginate(&allSales, 2, 2)
	// fmt.Println("Select Limit 2 Offset 2:")
	// for _, sale := range allSales {
	// 	fmt.Println(sale)
	// }

	// // Select, Where
	// db.Where("price > ?", 500).Find(&allSales)
	// fmt.Println("Select Where price > 500:")

	// for _, sale := range allSales {
	// 	fmt.Println(sale)
	// }

	// // Select, Where, Like
	// db.Where("description LIKE ?", "%Mouse%").Find(&allSales)
	// fmt.Println("Select Where description LIKE %Mouse%:")

	// for _, sale := range allSales {
	// 	fmt.Println(sale)
	// }

	// // Update
	// var s Sale
	// ID := 2
	// updateDescription(&s, ID, "New description")
	// db.First(&s, ID)
	// fmt.Println("Updated:")
	// fmt.Println(s)

	// // Delete
	// var s2 Sale
	// ID = 3
	// remove(&s2, ID)
}

func create(sale *Sale) {
	db.Create(sale)
}

func createMany(sales *[]Sale) {
	db.Create(&sales)
}

func findByDescription(sale *Sale, description string) {
	db.First(sale, "description = ?", description)
}

func findAll(sales *[]Sale) {
	db.Find(&sales)
}

func paginate(sales *[]Sale, limit int, offset int) {
	db.Limit(limit).Offset(offset).Find(sales)
}

func updateDescription(sale *Sale, ID int, description string) {
	db.First(sale, ID)
	sale.Description = description
	db.Save(sale)
}

func remove(sale *Sale, ID int) {
	db.First(sale, ID)
	db.Delete(sale)
}
