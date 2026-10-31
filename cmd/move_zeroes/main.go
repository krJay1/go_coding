package main

import "fmt"

func main() {
	var arr []int = []int{0, 1, 3, 0, 0, 5, 0, 7}
	sortZero(arr)
	fmt.Println(arr)
}

func sortZero(arr []int) {
	// idx := 0
	var ids []int

	for i := 0; i < len(arr); i++ {
		item := arr[i]
		if item == 0 {
			ids = append(ids, i)
			continue
		}
		if item > 0 {
			idx, ok := First(ids)
			if !ok {
				continue
			}
			arr[idx], arr[i] = arr[i], arr[idx]
			ids = append(ids, i)
			ids = ids[1:]
		}
	}
}

func First[T any](arr []T) (T, bool) {
	arrLen := len(arr)
	if arrLen > 0 {
		return arr[0], true
	}
	var Zero T
	return Zero, false
}
