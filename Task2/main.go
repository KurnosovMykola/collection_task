package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string

	strs := strings.Split(input, " ")
	var ints []int
	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}
	}

	//fmt.Println(ints)

	min := ints[0]
	max := ints[0]
	for _, element := range ints {
		if element > max {
			max = element
		}
		if element < min {
			min = element
		}
	}

	result = fmt.Sprintf("%d %d", max, min)
	fmt.Println(result)

}

//
