package main

import "fmt"

func div(dividen int, divisor int) int {

	i := 0
	x := 0

	for x < dividen {
		x = i * divisor
		i++
	}
	if x == dividen {
		return i - 1
	}

	return 0

}

func main() {

	x := div(20, 4)
	fmt.Println(x)

}
