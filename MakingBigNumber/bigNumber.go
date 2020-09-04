package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputNumbers() []int {
	slice := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, value := range strings.Split(scanner.Text(), " ") {
		toInt, _ := strconv.Atoi(value)
		slice = append(slice, toInt)
	}
	return slice
}

func main() {
	fmt.Println("A program to make a big number from a series of numbers..")
	fmt.Println("Input the numbers here")

	//taking an input of slice of numbers
	numberSlice := inputNumbers()

	fmt.Println(numberSlice)
}
