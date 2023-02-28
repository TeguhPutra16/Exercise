package main

import "fmt"

func main() {

	x := 10
	for x < 9 {
		x = x + 1
		fmt.Println(x)
	}

	fmt.Println(x)

}
