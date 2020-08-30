package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d\n", &N)
	var arr = [100000]int{}
	var truth = [100000]bool{}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
		if arr[i] == 0 {
			truth[i] = true
		} else {
			truth[arr[i]-1] = true
		}
	}
	for i := 0; i < N; i++ {
		if truth[i] == false {
			fmt.Printf("%d ", i+1)
		}
	}
}
