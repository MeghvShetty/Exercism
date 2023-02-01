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
	return make(map[string]int)
}

// Maps key passing
func Keypassing(val map[string]int, unit string) bool {
	u, ok := val[unit]
	fmt.Println(u, ok)
	return true
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	return true
}

// // RemoveItem removes an item from customer bill.
// func RemoveItem(bill, units map[string]int, item, unit string) bool {
// 	panic("Please implement the RemoveItem() function")
// }

// // GetItem returns the quantity of an item that the customer has in his/her bill.
// func GetItem(bill map[string]int, item string) (int, bool) {
// 	panic("Please implement the GetItem() function")
// }

func main() {
	fmt.Println(Units())
	fmt.Println(NewBill())
	// bill := NewBill()
	units := Units()
	data := Keypassing(units, "carrot ")
	fmt.Println(data)
	// ok := AddItem(bill, units, "carrot", "dozen")
}
return ok,false
}

return u, true