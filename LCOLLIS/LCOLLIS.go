package main

import (
	"fmt"
	"strconv"
)

func main() {
	var T, N, M, toAdd, result int
	var pr string

	fmt.Scan(&T)
	for T > 0 {
		result = 0
		fmt.Scanf("%d %d", &N, &M)
		// fmt.Println(T, N, M)

		arr := make([]int, M)

		//To add all the possible collisions
		for i := 0; i < N; i++ {
			fmt.Scan(&pr)
			for j := 0; j < M; j++ {
				toAdd, _ = strconv.Atoi(string(pr[j]))
				arr[j] += toAdd
			}
			//fmt.Println(pr)
		}
		//total collisions
		for i := 0; i < M; i++ {
			result += (arr[i] * (arr[i] - 1)) / 2
		}
		fmt.Println(result)
		T = T - 1
	}
}
