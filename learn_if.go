package main

import "fmt"

func main() {
	for b := 0; b < 20; b++ {
		if b > 10 {
			fmt.Print("a>0")
			fmt.Println("if条件不要求用括号")
		} else {
			fmt.Print("a<=10")
		}
	}
}
