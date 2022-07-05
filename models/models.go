package models

import "fmt"

func init() {
	fmt.Println("Models")
}

type Person struct {
	Name     string
	Balance  float64
	Currency string
}

// Create initial map of customers to be used in the program
func Customers() map[int]*Person {
	// maps hold references to their values so do not need pointers in methods: https://go.dev/blog/maps
	// maps are unordered, so we need to sort them
	customers := map[int]*Person{
		456: {
			Name:     "John",
			Balance:  100.00,
			Currency: "$",
		},
		123: {
			Name:     "Jane",
			Balance:  200.00,
			Currency: "£",
		},
		789: {
			Name:     "Joe",
			Balance:  300.00,
			Currency: "€",
		},
	}
	return customers
}

func (p Person) CheckBalance() string {
	return fmt.Sprintf("%s has %s%.2f", p.Name, p.Currency, p.Balance)
}

func (p *Person) Withdraw(amount float64) string {
	p.Balance -= amount
	return p.CheckBalance()
}

func (p *Person) Deposit(amount float64) string {
	p.Balance += amount
	return p.CheckBalance()
}
