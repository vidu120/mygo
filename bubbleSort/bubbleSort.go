package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//function for taking input and converting it to []int format
func inputSlice(OriginalSlice []int) []int {
	//reader for standard input
	reader := bufio.NewReader(os.Stdin)
	inStr, _ := reader.ReadString('\n')

	//parsing through the string and extracting int values from it
	for _, value := range strings.Split(strings.TrimRight(inStr, "\n"), " ") {
		intConversion, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}
		OriginalSlice = append(OriginalSlice, intConversion)
	}

	//returning the []int slice
	return OriginalSlice
}

//for swapping two adjacent places in a slice under a given condition
func swap(swap []int, index int) {
	if swap[index] > swap[index+1] {
		store := swap[index+1]
		swap[index+1] = swap[index]
		swap[index] = store
	}
}

//using the bubble sort concept to sort the slice of any length
func bubbleSort(toSort []int) {
	for i := 0; i < len(toSort); i++ {
		for j := 0; j < len(toSort)-i-1; j++ {
			swap(toSort, j)
		}
	}
}

func main() {

	//prompt for input
	fmt.Println("Please input the set of integers in the format - (2 3 5..)")

	//declaring the slice
	originalSlice := make([]int, 0, 10)

	//took the input as a string and converted it to []int format
	originalSlice = inputSlice(originalSlice)
	// fmt.Println(originalSlice)

	//sorting using bubble sort in the original slice itself
	bubbleSort(originalSlice)

	//printing the bubble sorted slice
	fmt.Print("Bubble Sorted Slice in ascending order :\n", originalSlice, "\n")
}
