package main

import "fmt"

func main() {
	var T, a, b, c, x, y int
	fmt.Scan(&T)
	for T > 0 {
		fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &x, &y)

		//making an array
		var arr [3]int = [3]int{a, b, c}

		var ifTrue bool = false

		for i := 0; i < 3 && !ifTrue; i++ {
			for j := 0; j < 3 && !ifTrue; j++ {
				if i == j || x-arr[i] < 0 || y-arr[j] < 0 {
					continue
				} else if x+y-arr[i]-arr[j] == arr[3-i-j] {
					ifTrue = true
				}
			}
		}
		if ifTrue {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		T = T - 1
	}
}
