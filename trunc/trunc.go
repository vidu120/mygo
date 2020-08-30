package main

import "fmt"

func main() {
	fmt.Print("Enter a floating point number to be truncated - \n")

	var numberToBeTruncated float64

	fmt.Scan(&numberToBeTruncated)

	var truncatedNumber = int64(numberToBeTruncated)

	fmt.Printf("Therefore the truncated form of the number %f is %d", numberToBeTruncated, truncatedNumber)
}
