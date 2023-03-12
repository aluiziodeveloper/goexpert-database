package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Sale struct {
	ID    string
	Name  string
	Price float64
}

func NewSale(name string, price float64) *Sale {
	return &Sale{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Insert
	sale := NewSale("Notebook", 2990.90)
	err = insertSale(db, *sale)
	if err != nil {
		panic(err)
	}
	// Update
	sale.Name = "MacBook"
	sale.Price = 19099.89
	err = updateSale(db, sale)
	if err != nil {
		panic(err)
	}
	// Select One
	s, err := selectSale(db, sale.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Produto: %v. Preço: R$ %.2f.\n", s.Name, s.Price)
	// Select All
	sales, err := selectAllSales(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Lista de Produtos")
	for _, s := range sales {
		fmt.Printf("Produto: %v. Preço: R$ %.2f.\n", s.Name, s.Price)
	}
	// Delete
	err = deleteSale(db, sale.ID)
	if err != nil {
		panic(err)
	}
}

func insertSale(db *sql.DB, sale Sale) error {
	stmt, err := db.Prepare("insert into sales(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sale.ID, sale.Name, sale.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateSale(db *sql.DB, sale *Sale) error {
	stmt, err := db.Prepare("update sales set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sale.Name, sale.Price, sale.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectSale(db *sql.DB, id string) (*Sale, error) {
	stmt, err := db.Prepare("select id, name, price from sales where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var s Sale
	err = stmt.QueryRow(id).Scan(&s.ID, &s.Name, &s.Price)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func selectAllSales(db *sql.DB) ([]Sale, error) {
	rows, err := db.Query("select id, name, price from sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sales []Sale
	for rows.Next() {
		var s Sale
		err = rows.Scan(&s.ID, &s.Name, &s.Price)
		if err != nil {
			return nil, err
		}
		sales = append(sales, s)
	}
	return sales, nil
}

func deleteSale(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from sales where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
