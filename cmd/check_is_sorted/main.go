package main

import "fmt"

func main() {
	var arr []int = []int{1, 22, 333, 55}
	fmt.Println(checkSorted(arr))

}

func checkSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
