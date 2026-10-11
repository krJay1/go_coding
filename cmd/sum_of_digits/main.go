package main

import "fmt"

// 01. Sum of Digits
// Input: An integer N (e.g., 1234)
// Output: Sum of its digits (e.g., 1 + 2 + 3 + 4 = 10)
// Constraint: Do not convert the integer to a string. Use '%' and '/'.
func main() {
	fmt.Println(getIntSum(8437464290))
}

func getLastDigit(n int) (int, int) {
	var quotient int = n / 10
	var rem int = n % 10
	return quotient, rem
}

func getIntSum(n int) int {

	var sum int = 0
	var num int = n
	for {
		q, r := getLastDigit(num)
		fmt.Println(q, r)
		sum += r
		num = q
		if q < 10 {
			sum += q
			break
		}
	}
	return sum
}
