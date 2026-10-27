package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{2, 3, 5, 66, 7}
	start := time.Now()
	// idx1, idx2 := getIndices(arr, 7)
	idx1, idx2 := getOptimizedIdxs(arr, 5)
	fmt.Println(arr[idx1], arr[idx2], time.Since(start))

}

// func getIndices(arr []int, target int) (int, int) {
// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(arr); j++ {
// 			if target == arr[i]+arr[j] {
// 				return i, j
// 			}
// 		}
// 	}
// 	return -1, -1
// }

func getOptimizedIdxs(arr []int, target int) (int, int) {
	seen := map[int]int{}
	for i, num := range arr {
		complement := target - num
		if idx, ok := seen[complement]; ok {
			return i, idx
		}
		seen[num] = i
	}

	return -1, -1
}
