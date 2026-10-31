package main

import "fmt"

func main() {
	var arr []int = []int{0, 1, 3, 0, 0, 5, 0, 7}
	sortZero(arr)
	fmt.Println(arr)
}

func sortZero(arr []int) {
	// idx := 0
	var initial int = 0

	for i := 0; i < len(arr); i++ {
		item := arr[i]
		if item != 0 {
			arr[initial], arr[i] = arr[i], arr[initial]
			initial++
		}
	}
}
