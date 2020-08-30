package main

import "fmt"

func main() {
	var n, i int
	fmt.Scanf("%d\n", &n)
	for i < n {
		var N int
		fmt.Scanf("%d\n", &N)
		var pr string
		fmt.Scanf("%s\n", &pr)
		var code = []byte(pr)
		for j := 1; j < N; j = j + 2 {
			code[j], code[j-1] = code[j-1], code[j]
			code[j-1] = byte(219 - int(code[j-1]))
			code[j] = byte(219 - int(code[j]))
		}
		if N%2 != 0 {
			code[N-1] = byte(219 - int(code[N-1]))
		}
		for j := 0; j < N; j++ {
			fmt.Printf("%c", code[j])
		}
		fmt.Print("\n")
		i++
	}
}
