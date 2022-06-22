package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	for i := 0; i < len(arr); i++ {
		Duplicate := false
		for x := 0; x < len(result); x++ {
			if arr[i] == result[x] {
				Duplicate = true
				break
			}
		}
		if !Duplicate {
			result = append(result, arr[i])
		}
	}
	fmt.Println(result)
}
