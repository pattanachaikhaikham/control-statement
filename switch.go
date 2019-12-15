package main

import "fmt"

func main() {
	fmt.Printf("impot(number) =")
	var num int32
	fmt.Scan(&num)
	switch num {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("unknow")
	}
}
