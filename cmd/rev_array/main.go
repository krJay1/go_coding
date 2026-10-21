package main

import "fmt"

func main() {
	arr := []int{2, 44, 5, 23, 65}
	revArr(arr)
	fmt.Println(arr)

}

func revArr(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
