package main

import "fmt"

func main() {

	//Let's make a slice
	//1st Way
	slic := []int{1, 2, 3, 4}
	fmt.Print(slic, len(slic), cap(slic))
	slic = append(slic, 5)
	fmt.Print(slic, len(slic), cap(slic))

	slicIt := slic[0:2]
	fmt.Print(slicIt, len(slicIt), cap(slicIt))
	slicIt[0] = 9
	slicIt[1] = 10
	fmt.Print(slic, len(slic), cap(slic))

	//2nd way
	slic1 := [...]int{3, 45, 6, 7, 4}
	fmt.Print(slic1[1:3], len(slic1), cap(slic1))

	//3rd way
	slic2 := make([]int, 10, 100)
	fmt.Print(slic2, len(slic2), cap(slic2))
	slic2 = append(slic2, 19)
	fmt.Print(slic2, len(slic2), cap(slic2))

	//New thing i learned
	fmt.Println()
	mine := [3]int{1, 2, 3}
	for i, j := range mine {
		fmt.Println(i, j)
	}

	//hash tables
	var mine0 map[string]int
	mine0 = map[string]int{"joe": 1, "jane": 2, "Vidhan": 5}
	fmt.Print(mine0)

	var mine1 map[string]int
	mine1 = make(map[string]int)
	fmt.Print(mine1, len(mine1))

	mine3 := make(map[string]int)
	fmt.Print(mine3, len(mine3))
}
