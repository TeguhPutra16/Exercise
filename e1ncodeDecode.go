package main

import (
	"fmt"
	"strings"
)

func encodeDecode(word string) string {
	if strings.HasPrefix(word, "<encode>") {

		word = strings.TrimPrefix(word, "<encode>")

		encodeWord := ""

		for _, v := range word {
			if v >= 97 && v <= 122 {
				v = v + 2

				if v > 122 {
					v = v - 26
				}
			}
			encodeWord += string(v)
		}
		return encodeWord
	} else if strings.HasPrefix(word, "<decode>") {
		word = strings.TrimPrefix(word, "<decode>")

		decodeWord := ""

		for _, v := range word {
			if v >= 97 && v <= 122 {
				v = v - 2

				if v < 97 {
					v = v + 26
				}
			}
			decodeWord += string(v)
		}
		return decodeWord
	}
	return ""
}

func main() {
	fmt.Println(encodeDecode("<encode>abcd")) //cdef
	fmt.Println(encodeDecode("<encode>xyz"))  //zab

	fmt.Println(encodeDecode("<decode>cdef")) //abcd
	fmt.Println(encodeDecode("<decode>zab"))  //xyz
}
