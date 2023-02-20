package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := ""
	for i := 0; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			str1 := strconv.Itoa(i)
			res += str1 + "FizzBuzz"
		} else if i%3 == 0 {
			str2 := strconv.Itoa(i)
			res += str2 + "Fizz"
		} else if i%5 == 0 {
			str3 := strconv.Itoa(i)
			res += str3 + "Buzz"

		} else {
			continue
		}

	}
	fmt.Println(res)

}
