package main

import "fmt"

func main() {

	s := "MCMXCIV"

	x := romanToInt(s)
	fmt.Println(x)

}

func romanToInt(s string) (number int) {
	symbolMap := makeRomanSymbolMap()
	// sArray := []rune(s)
	// fmt.Println("sArray", sArray)
	for i := 0; i < len(s); i++ {
		// fmt.Println("sArray i", sArray[i])

		if i+1 >= len(s) {
			number += symbolMap[string(s[i])]
		} else if symbolMap[string(s[i])] < symbolMap[string(s[i+1])] {
			number -= symbolMap[string(s[i])]
		} else {
			number += symbolMap[string(s[i])]
		}
	}
	return number
}

func makeRomanSymbolMap() map[string]int {
	symbolMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	return symbolMap
}
