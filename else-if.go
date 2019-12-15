package main

import "fmt"

func main() {
	fmt.Printf("impot = ")
	var score int32
	fmt.Scan(&score)
	if score >= 80 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B+")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 65 {
		fmt.Println("C+")
	} else if score >= 60 {
		fmt.Println("C")
	} else if score >= 55 {
		fmt.Println("D+")
	} else if score >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
