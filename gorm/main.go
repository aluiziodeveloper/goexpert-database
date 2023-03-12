package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Sale struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
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
		Name:  "Product 01",
		Price: 1200,
	})

	// Update
	var s Sale
	ID := 1
	updateName(&s, ID, "New name")

	// Delete
	var s2 Sale
	remove(&s2, ID)

	// // Create Many
	// sales := []Sale{
	// 	{Name: "Notebook dell", Price: 3929.90},
	// 	{Name: "MacBook Pro", Price: 9129.90},
	// 	{Name: "Monitor", Price: 829.90},
	// 	{Name: "Keyboard", Price: 29.90},
	// 	{Name: "Mouse", Price: 29.90},
	// }
	// createMany(&sales)

	// // Select One
	// var sale Sale
	// findByName(&sale, "Monitor")
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
	// db.Where("name LIKE ?", "%Mouse%").Find(&allSales)
	// fmt.Println("Select Where name LIKE %Mouse%:")

	// for _, sale := range allSales {
	// 	fmt.Println(sale)
	// }

	// // Update
	// var s Sale
	// ID := 2
	// updateName(&s, ID, "New name")
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

func findByName(sale *Sale, name string) {
	db.First(sale, "name = ?", name)
}

func findAll(sales *[]Sale) {
	db.Find(&sales)
}

func paginate(sales *[]Sale, limit int, offset int) {
	db.Limit(limit).Offset(offset).Find(sales)
}

func updateName(sale *Sale, ID int, name string) {
	db.First(sale, ID)
	sale.Name = name
	db.Save(sale)
}

func remove(sale *Sale, ID int) {
	db.First(sale, ID)
	db.Delete(sale)
}
