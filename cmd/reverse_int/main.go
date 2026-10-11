package main

import (
	"fmt"
	"strconv"
)

// 02. Reverse an Integer
// Input: An integer x (e.g., 123, -456)
// Output: Reversed integer (e.g., 321, -654)
// Constraint: Handle negative numbers properly without using strings.
func main() {
	fmt.Print(revInt(1234))

}

func revInt(n int) int {
	num := n
	rev := ""

	for {
		l := num % 10
		rev += strconv.Itoa(l)
		num = num / 10
		if num < 10 {
			rev += strconv.Itoa(num)
			break
		}

	}
	number, _ := strconv.ParseInt(rev, 10, 36)
	return int(number)

}
