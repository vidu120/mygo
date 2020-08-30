package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hey there ,please input integers one by one to sort them..")
	mainSlice := make([]int, 0, 100)
	input := ""
	toAdd := 0
	var once bool = true
	for {
		fmt.Println("Please enter an integer or X to exit: ")
		fmt.Scan(&input)
		if input == "X" || input == "x" {
			fmt.Println("Quitting")
			break
		}
		toAdd, _ = strconv.Atoi(input)
		if once {
			mainSlice = append(mainSlice, toAdd)
			once = false
			fmt.Print("Sorted: ", mainSlice, "\n")
			continue
		}
		for i := 0; i < len(mainSlice); i++ {
			if mainSlice[i] > toAdd {
				mainSlice = append(mainSlice, 0)
				src := mainSlice[i+1 : len(mainSlice)]
				dst := mainSlice[i : len(mainSlice)-1]
				copy(src, dst)
				mainSlice[i] = toAdd
				break
			}
			if i == len(mainSlice)-1 {
				mainSlice = append(mainSlice, toAdd)
				break
			}
		}
		fmt.Print("Sorted: ", mainSlice, "\n")
	}
}
