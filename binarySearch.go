package main

import (
	"fmt"
	"math/rand"
	"time"
)


func quickSort(digits []int) []int {
	if len(digits) < 2 {
		return digits
	}
	if len(digits) == 2 {
		if digits[0] > digits[1] {
			digits[0], digits[1] = digits[1], digits[0]
		}
		return digits
	}
	base := digits[0]
	digits = digits[1:]
	var left, right []int
	for _, elem := range digits {
		if elem <= base {
			left = append(left, elem)
		} else {
			right = append(right, elem)
		}
	}
	left = append(left, base)
	return append(quickSort(left), quickSort(right)...)
}

func binarySearch(arr[]int, x int) (int) {
	l, m:= 0, 0
	r := len(arr) + 1
	for l < r - 1 {
		m = (l + r) / 2
		if (arr[m] <= x){
			l = m
		} else {
			r = m
		}
	}
	return r
}

func main() {
	var arr []int
	var x int
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10; i++ {
		arr = append(arr, int(rand.Int31()%100))
	}
	arr = quickSort(arr)
	x = int(rand.Int31()%100)
	fmt.Println(arr)
	fmt.Println(x)
	fmt.Printf("place for x is: %d\n", binarySearch(arr, x))
}
