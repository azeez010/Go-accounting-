package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "time"
	// "fmt"
  )
  
func Open() (*gorm.DB, error) {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("account.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		}

		// Migrate the schema
		db.AutoMigrate(&Expenses{}, &Product{}, &Sales{})
		// Create
		// db.Create(&Product{Name: "Rice", CostPrice: 17000, SellingPrice: 23500, Quantity: 5, Date: time.Now()})


		 // Read
		//  var product Product

		//  db.First(&product, 1) // find product with integer primary key

		//  db.First(&product, "Name = ?", "Rice") // find product with code D42
	   
		// fmt.Println(product)

		return db, nil
} 
  