package main

import "fmt"

type person struct {
	Nama string
	Umur int
}

func main() {

	var orang person = person{"teguh", 17}

	var orang1 *person = &orang

	*orang1 = person{"asu", 90}

	fmt.Println(orang)
	fmt.Println(orang1)

}
