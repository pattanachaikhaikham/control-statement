package main

import "fmt"

func main() {
	for num := 1; num <= 1000; num = num + 1 {
		if num%3 == 0 {
			fmt.Println(num, "fizz")
		} else if num%5 == 0 {
			fmt.Println(num, "buzz")
		} else {
			fmt.Println(num)
		}
	}
}
