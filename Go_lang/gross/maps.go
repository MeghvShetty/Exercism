package main

import (
	"fmt"
)

func main() {
	// way 1 to create a map
	var mp map[string]int = map[string]int{
		"apple":  1,
		"banana": 6,
		"orange": 9,
	}
	// way 2 to create a map
	//mp1 := make(map[string]int)
	fmt.Println(mp)
	fmt.Println(mp["apple"])
	// add value to map
	mp["pear"] = 50
	fmt.Println(mp)
	// changing a value
	mp["apple"] = 14
	fmt.Println(mp)
	// delete a key in the map
	delete(mp, "apple")
	fmt.Println(mp)
}
