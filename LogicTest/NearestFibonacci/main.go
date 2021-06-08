package main

import "fmt"

func nearestFibonacci(sum int) int {
	first, second := 0, 1
	third := first + second

	for third <= sum {
		first = second
		second = third
		third = first + second
	}

	result := third - sum
	return result
}

func main() {
	arr := []int{15, 1, 3}
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	result := nearestFibonacci(sum)
	fmt.Println(result)
}
