package main

import "fmt"

func main() {
	for num := 1; num <= 10; num = num + 1 {
		if num%2 == 0 {
			fmt.Println(num, "even Number")
		} else {
			fmt.Println(num, "odd Number")
		}
	}
}
