package main

import (
	"fmt"
)

func main() {
	fmt.Println("Калькулятор транзакций")
	transactions := [5]int{1, 2, 3, 4, 5}
	banks := [2]string{}

	fmt.Println(transactions[1])
	banks[0] = "VTB"
	fmt.Println(banks)
	partial := transactions[1:4]
	fmt.Println(partial)
}
