package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var units map[string]int = map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	i := 0
	for key, Value := range units {
		fmt.Println()

	}
	return false
}

func main() {
	bill := NewBill()
	units := Units()
	ok := AddItem(bill, units, "carrot", "test")
	fmt.Println(ok)
	fmt.Println(bill)

}
