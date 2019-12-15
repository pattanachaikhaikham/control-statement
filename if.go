package main

import "fmt"

func main() {
	num := 0
	for {
		fmt.Println(num)
		num = num + 1
		if num > 10 {
			break
		}
	}
}
