package database

import (
	"gorm.io/gorm"
	"time"
  )


// type ExpenseDuration string

// const (
// 	Daily ExpenseDuration = "daily"
// 	Weekly ExpenseDuration = "weekly"
// 	Monthly ExpenseDuration = "monthly"
// 	Yearly ExpenseDuration = "yearly"

// )

// func (e *ExpenseDuration) Scan(value interface{}) error {
// 	*e = ExpenseDuration(value.([]byte))
// 	return nil
// }

// func (e ExpenseDuration) Value() (driver.Value, error) {
// 	return string(e), nil

type Product struct {
	gorm.Model
	Name  string
	CostPrice uint
	SellingPrice uint
	Quantity uint
	Date time.Time
}

type Expenses struct {
	gorm.Model
	Name string
	Duration string 
	Date time.Time
	Cost uint
}

// `gorm:"type:enum('daily', 'monthly', 'weekly', 'yearly');default:'monthly'"`
type Sales struct {
	gorm.Model
	Name  string
	TotalSales uint
	Profit uint
	Quantity uint
	Date time.Time
}
