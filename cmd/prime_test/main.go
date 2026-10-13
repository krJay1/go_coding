package main

import "fmt"

func main() {
	fmt.Println("Check Prime:", checkPrime(20))

}

func checkPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		rem := n % i
		fmt.Println(i, rem)
		if rem == 0 && i != n {
			return false
		}
	}
	return true

}
