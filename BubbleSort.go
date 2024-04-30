package main

import "fmt"

func main() {

	var array = []uint{1, 5, 9, 100, 69, 42, 76, 55, 900, 7, 4, 25, 108}

	fmt.Printf("Sorted array is: %v", BubbleSort(array))
}

func BubbleSort(numbers []uint) []uint {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				var tmp = numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = tmp
			}
		}
	}
	return numbers
}
