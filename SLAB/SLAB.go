package main

import "fmt"

func main() {

	var T, N int32
	fmt.Scan(&T)
	for T > 0 {
		fmt.Scan(&N)

		if N == 0 {
			fmt.Println(0)
			T = T - 1
			continue
		}
		var ifTrue bool = N%250000 == 0
		var toMultiply int32 = N / 250000

		if ifTrue {
			toMultiply--
		}

		if toMultiply >= 6 {
			fmt.Println(N - int32(187500.0+float32((N-1500000)*3/10)))
			T = T - 1
			continue
		}

		//setting the tax percent
		var tax float32 = float32((toMultiply*(toMultiply-1))/2) * 0.05 * 250000.0
		tax += float32(N-toMultiply*250000) * float32(toMultiply) * 0.05
		fmt.Println(N - int32(tax))
		T = T - 1
	}
}
