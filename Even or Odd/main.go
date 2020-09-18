package main

import (
	"strconv"
)

func main() {
	checkEvenOrOdd([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func checkEvenOrOdd(numbers []int) string {
	eval := ""

	for _, num := range numbers {
		rem := num % 2

		if rem != 0 {
			//fmt.Printf("%v is odd", num)
			// res = num + "is odd"
			eval = strconv.FormatInt(int64(num), 10) + " is odd"
		} else {
			// fmt.Printf("%v is even", num)
			eval = strconv.FormatInt(int64(num), 10) + " is even"
		}
	}

	return eval
}
