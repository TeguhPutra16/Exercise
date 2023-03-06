package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7}
	seq := []int{1, 3, 5, 7}

	fmt.Println(Sequence(array, seq))

}

func Sequence(array []int, seq []int) bool {

	lengtArr := len(array)
	j := 0

	for i := 0; i < lengtArr; i++ {
		fmt.Println(len(array), "array", array)
		// fmt.Println(j)
		fmt.Println(i, "i")
		// fmt.Println(array[i])
		if i == len(array) {
			return len(array) == len(seq)
		}
		if array[i] != seq[j] {
			array = append(array[:i], array[i+1:]...)
			i--

		} else {
			j++

		}

		// fmt.Println(j)
		if len(seq) == j {
			// fmt.Println((array), len(seq))
			return len(array) == len(seq)
		}

	}
	return false

}
