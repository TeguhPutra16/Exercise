package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	strs := []string{"cir", "car"}

	fmt.Println(sufix(strs))

}

func sufix(strs []string) string {
	sort.SliceStable(strs, func(i, j int) bool {
		return strs[i] > strs[j]
	})
	suff := ""
	if len(strs) == 1 {
		suff = strs[0]
		return suff
	}
	for k, v := range strs[0] {

		for _, str := range strs {
			fmt.Println("suff1", suff)
			fmt.Println("str", str)
			if !strings.HasPrefix(str, suff) {
				fmt.Println("suff2", suff)
				return suff[:len(suff)-1]
			} else if k == len(strs[0])-1 {
				return suff

			}

		}
		suff += string(v)

	}
	return suff

}
func longestCommonPrefix(strs []string) string {
	suff := ""
	for i, ch := range strs[0] {

		fmt.Println("i", i)
		fmt.Println("ch", ch)
		for _, str := range strs {

			if i > len(str)-1 {
				fmt.Println("oke1")
				return suff
			} else if rune(str[i]) != ch {
				fmt.Println("oke2")
				return suff
			}
		}

		suff += string(ch)
	}
	return suff

}
