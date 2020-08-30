package main

import "fmt"

func main() {
	var n, i int
	fmt.Scanf("%d\n", &n)
	for i < n {
		var length int
		fmt.Scanf("%d\n", &length)
		var line, line0 string
		fmt.Scanf("%s\n", &line)
		fmt.Scanf("%s\n", &line0)
		var pr int
		var pr0 int
		for j := 0; j < length; j++ {
			if line[j] == '1' {
				pr++
			}
			if line0[j] == '1' {
				pr0++
			}
		}
		if pr == pr0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		i++
	}
}
