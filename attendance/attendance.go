package main

import "fmt"

func main() {
	var n, i int
	fmt.Scanf("%d\n", &n)
	for i < n {
		var N int
		fmt.Scanf("%d\n", &N)
		var name = [100][2]string{}
		for j := 0; j < N; j++ {
			fmt.Scanf("%s", &name[j][0])
			fmt.Scanf("%s\n", &name[j][1])
		}
		var pr bool
		for j := 0; j < N; j++ {
			pr = false
			for k := 0; k < N; k++ {
				if j == k {
					continue
				}
				if name[j][0] == name[k][0] {
					pr = true
					break
				}
			}
			if pr == true {
				fmt.Printf("%s %s\n", name[j][0], name[j][1])
			} else {
				fmt.Printf("%s\n", name[j][0])
			}
		}
		i++
	}
}
