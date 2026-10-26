package main

import "fmt"

func main() {
	var arr []int = []int{1, 1, 1, 1, 2, 3, 3, 3, 4, 4, 5, 5, 6, 7, 7, 8}
	count, err := countK(arr)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(count)
}

func countK(arr []int) (int, error) {
	idx := 1
	var err error
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			continue
		}
		if arr[i] > arr[i+1] {
			idx = 0
			err = fmt.Errorf("Not Sorted")
			break
		}

		if arr[i] < arr[i+1] {
			arr[idx] = arr[i]
			idx++
		}
	}

	return idx, err
}
