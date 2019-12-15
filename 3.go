package main

import "fmt"

func main() {
	for num := 1; num <= 1000; num = num + 1 {
		if num%2 == 0 {
			fmt.Print()
		} else if num%3 == 0 {
			fmt.Println(num)
		} else {
			fmt.Print()
		}
	}
}
