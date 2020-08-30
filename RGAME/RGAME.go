package main

import (
	"fmt"
)

func main() {
	var T int
	var N int
	divider := 1000000007
	fmt.Scanf("%d\n", &T)
	for T > 0 {
		fmt.Scanf("%d\n", &N)
		score := 0
		elem := 0
		fmt.Scanf("%d", &elem)
		ty := elem
		mul := 1
		for i := 1; i < N+1; i++ {
			fmt.Scanf("%d", &elem)
			score = (2*score + elem*ty) % divider
			ty = (ty + mul*elem) % divider
			mul = mul * 2
		}
		fmt.Println(score)
		T = T - 1
	}
}
