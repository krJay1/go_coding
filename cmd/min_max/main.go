package main

import "fmt"

func main() {
	arr := []int{7, 2, 44, 4, 5, 930, 4554, 222, 34, 33, 4, 6}
	min, max := getMinMax(arr)
	fmt.Println(min, max)
}

func getMinMax(arr []int) (int, int) {
	arrLen := len(arr)
	var min int
	var max int
	for i := 0; i < arrLen; i++ {
		item := arr[i]
		if i == 0 {
			min = item
		} else if item <= min {
			min = item
		}

		if item >= max {
			max = item
		}
	}
	return min, max

}
