package main

import "fmt"

func main() {
	var n, i int
	fmt.Scanf("%d\n", &n)
	var N, Q int
	var S string = "jdvn"
	for i < n {
		fmt.Scanf("%d\n%s\n%d\n", &N, &S, &Q)
		mine := []rune(S)
		var a, b, c, d int
		var pr0 int
		pr := 'a'
		for j := 0; j < Q; j++ {
			fmt.Scanf("%d ", &pr0)
			if pr0 == 1 {
				fmt.Scanf("%d %d %c\n", &a, &b, &pr)
				for j := a - 1; j < b-1; j++ {
					mine[j] = pr
				}

			} else {
				fmt.Scanf("%d %d %d %d\n", &a, &b, &c, &d)
			}
		}
		i++
	}
}
