package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(swap("Data1", "Data2"))
}
