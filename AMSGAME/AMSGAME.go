package main

import "fmt"

func gcd(x, y int) int {
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}

func main() {
	var t, n, mine int
	fmt.Scanf("%d\n", &t)
	for t > 0 {
		fmt.Scanf("%d\n", &n)
		var arr = [1000]int{}
		var i int
		fmt.Scanf("%d ", &arr[0])
		for i := 1; i < n; i++ {
			fmt.Scanf("%d ", &arr[i])
			mine = gcd(min())
		}
		t = t - 1
	}
}
