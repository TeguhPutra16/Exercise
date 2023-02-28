package main

import "fmt"

/*
    *
   * *
  * * *
 * * * *
* * * * *


*/

func main() {
	res := ""
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print(" ")

		}
		res = res + "* "
		fmt.Print(res)
		fmt.Println()

	}

}
