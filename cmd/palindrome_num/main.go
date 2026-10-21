package main

import "fmt"

func main() {
	fmt.Println(isPalin(123121))

}

func isPalin(num int) bool {
	var temp int = num
	var rev int = 0
	var rem int
	for {
		rem = temp % 10
		temp = temp / 10
		rev = rev*10 + rem
		if temp < 10 {
			rev = rev*10 + temp
			break
		}
	}
	fmt.Println(rev)
	if rev == num {
		return true
	}
	return false
}
