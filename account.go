package main

import (
	"os"
	"fmt"
	"accounting/database"
	"accounting/controller"
	"gorm.io/gorm"
)

func main(){
	fmt.Println("Running...")
	db, err := database.Open()

	if err != nil {
		panic("Database Failed!")
		os.Exit(3)
	}

	Welcome(db)
}

func Welcome(db *gorm.DB) {
	var choice string;
	fmt.Println("Welcome to Our Accounting software!");
	
	fmt.Println("Choose the following \n1. Save Expenses \n2. Save Stock \n3. Save Sales \n4. Read Stock \n5. Calculate Expenses \n6. All Sales from Particular Date\n");
	fmt.Scanln(&choice)
	
	manageExpenses := controller.NewExpenses(db)
	switch choice {
		case "1":
			manageExpenses.SaveExpense()
		case "2":
			manageExpenses.SaveStock()
		case "3":
			manageExpenses.SaveSales()
		case "4":
			manageExpenses.ReadStock()
		case "6":
			manageExpenses.ReadSalesFrom()
	
	}
}
