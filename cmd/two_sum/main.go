package main

import "fmt"

func main() {
	arr := []int{2, 3, 5, 66, 7}
	idx1, idx2 := getIndices(arr, 7)
	fmt.Println(arr[idx1], arr[idx2])

}

func getIndices(arr []int, target int) (int, int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if target == arr[i]+arr[j] {
				return i, j
			}
		}
	}
	return 0, 0

}
