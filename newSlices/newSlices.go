package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//create an integer slice of size 3
	s := make([]int, 0, 3)

	fmt.Println("s: ", s)
	fmt.Println("len: ", len(s))
	fmt.Println("cap:", cap(s))

	//loop until â€˜Xâ€™ is entered

	var stringInput string

	var counter int = 0

	for counter < 3 {
		fmt.Println("Enter an integer! or X to exit.")
		fmt.Scanln(&stringInput)
		stringInput = strings.ToUpper(stringInput)

		if stringInput == "X" {
			fmt.Println("Goodbye")
			os.Exit(0)
		}

		i, err := strconv.Atoi(stringInput)
		if err != nil {
			fmt.Println("Enter an integer! or X to exit")
			fmt.Scanln(&stringInput)
			stringInput = strings.ToUpper(stringInput)
			continue
		}

		s = append(s, i)
		//		s[counter] = i
		fmt.Println("counter: %d, value: %d", counter, s[counter])

		//sort the slice and print the resulting slice
		sort.Ints(s)
		fmt.Println(s)

		counter++
	}

	fmt.Println("Enter an integer! or X to exit.")
	fmt.Scanln(&stringInput)
	stringInput = strings.ToUpper(stringInput)
	for stringInput != "X" {
		i, err := strconv.Atoi(stringInput)
		if err != nil {
			fmt.Println("Enter an integer! or X to exit")
			fmt.Scanln(&stringInput)
			stringInput = strings.ToUpper(stringInput)
			continue
		}

		s = append(s, i)

		//sort the slice and print the resulting slice
		sort.Ints(s)

		fmt.Println(s)

		fmt.Scanln(&stringInput)
		stringInput = strings.ToUpper(stringInput)
	}

	fmt.Println("Goodbye")

}
