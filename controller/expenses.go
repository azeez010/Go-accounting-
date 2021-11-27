package controller

import (
	"accounting/database"
	"accounting/utils"
	"fmt"
	"time"
	"gorm.io/gorm"
)

// type ExpenseDuration string;

// const (
// 	daily ExpenseDuration = "daily"
// 	weekly ExpenseDuration = "weekly"
// )

type Expenses struct {
	db *gorm.DB
} 

func NewExpenses(db *gorm.DB) *Expenses{
	return &Expenses{db}
}

func (exp *Expenses) SaveExpense(){
	var cost string
	var name string
	// var askDuration ExpenseDuration
	var duration string;
	
	var costInt uint

	fmt.Println("Enter Expenses COST\n")
	fmt.Scanln(&cost)
	
	defer func(){
		if err := recover(); err != nil{
			exp.SaveExpense()
		}
	}()

	utils.ToNumber(cost, &costInt)
	fmt.Println("Enter Expenses NAME\n")
	
	fmt.Scanln(&name)
	fmt.Println("Enter Expenses DURATION\n")
	
	fmt.Scanln(&duration)

	exp.db.Create(&database.Expenses{Cost: costInt, Name: name, Duration: duration})
}

func (exp *Expenses) SaveStock(){
	var costPrice string
	var sellingPrice string
	var name string
	// var askDuration ExpenseDuration
	var quantity string;
	
	var costPriceInt uint
	var sellingPriceInt uint
	var quantityInt uint
	

	fmt.Println("Enter Product COST price\n")
	fmt.Scanln(&costPrice)
	
	defer func(){
		if err := recover(); err != nil{
			exp.SaveStock()
		}
	}()

	utils.ToNumber(costPrice, &costPriceInt)
	fmt.Println("Enter Product NAME\n")
	
	fmt.Scanln(&name)
	fmt.Println("Enter Product Quantity\n")
	
	fmt.Scanln(&quantity)
	utils.ToNumber(quantity, &quantityInt)
	
	fmt.Println("Enter Product Selling Price\n")
	
	fmt.Scanln(&sellingPrice)
	utils.ToNumber(sellingPrice, &sellingPriceInt)
	
	if costPriceInt > sellingPriceInt {
		fmt.Println("Cost price must not be greater than selling price!")
		exp.SaveStock()
	} else {
		var product database.Product
		
		result := exp.db.Where("name = ?", name).First(&product)
		if result.RowsAffected > 0 {
			product.CostPrice =  costPriceInt
			product.SellingPrice = sellingPriceInt 
			product.Quantity = quantityInt 
			product.Name = name
			product.Date = time.Now()

			exp.db.Save(&product)
		} else {
			exp.db.Create(&database.Product{CostPrice: costPriceInt, 
				SellingPrice: sellingPriceInt, Quantity: quantityInt, Name: name,
				Date: time.Now(),
			})
		}
	}
}


func (exp *Expenses) SaveSales(){
	var name string
	var quantity string;
	
	var quantityInt uint
	

	
	defer func(){
		if err := recover(); err != nil{
			exp.SaveSales()
		}
	}()

	fmt.Println("Enter Sales NAME\n")
	
	fmt.Scanln(&name)
	fmt.Println("Enter Sales Quantity\n")
	
	fmt.Scanln(&quantity)
	utils.ToNumber(quantity, &quantityInt)
	
	var product database.Product
	allData := exp.db.First(&product, "name = ?", name )
	
	if allData.RowsAffected == 0 {
		fmt.Println("Product does exist!")
	} else {
		if product.Quantity < quantityInt {
			fmt.Println("Not enough product quantity!")
			exp.SaveSales()
		} else {
			// Update the stock quantity
			product.Quantity -= quantityInt
			exp.db.Save(&product)

			// Calculate Profit and Total Sales
			sales := product.SellingPrice * quantityInt
			var profit uint = (product.SellingPrice - product.CostPrice) * quantityInt
			
			// Save Sales
			exp.db.Create(&database.Sales{
				Quantity: quantityInt, 
				TotalSales: sales, 
				Name: product.Name,
				Profit: profit,
				Date: time.Now(),
			})
		}
	}
	
}

func (exp *Expenses) ReadStock(){
	var products []database.Product
	allData := exp.db.Find(&products)
	fmt.Println("Name", "Quantity", "Selling Price", "Cost Price", "Date")
	for i := 0; i < int(allData.RowsAffected); i++ {
		fmt.Printf("%v %v %v %v %v\n", products[i].Name, products[i].Quantity, products[i].SellingPrice, products[i].CostPrice, products[i].Date.Month() )
	}
}

func (exp *Expenses) ReadSalesFrom(date ...time.Time){
	if len(date) == 0 {
		var sales []database.Sales
		var totalSales uint
		var totalProfit uint

		allData := exp.db.Find(&sales)
		fmt.Println("Name", "Total Sales", "Profit", "Quantity", "Date")
		for i := 0; i < int(allData.RowsAffected); i++ {
			fmt.Printf("%v %v %v %v %v\n", sales[i].Name, sales[i].TotalSales, sales[i].Profit, sales[i].Quantity, sales[i].Date.Month() )
			totalSales += sales[i].TotalSales
			totalProfit += sales[i].Profit
		}

		fmt.Println("Total Sales: ", totalSales, "  ----  Total Profit: ", totalProfit)
	}
}
