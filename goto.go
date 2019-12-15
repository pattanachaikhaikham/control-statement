package main

import "fmt"

func main() {
	var num int
ReadInput:
	fmt.Print("type number :")
	fmt.Scan(&num)
	if num < 30 {
		goto ReadInput
	}
	fmt.Println(num)
}
