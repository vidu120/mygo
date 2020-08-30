package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter the String: ")
	var toTest string
	fmt.Scan(&toTest)

	//Below are the boolean conditions that we need to check
	ifPrefixI := strings.HasPrefix(strings.ToLower(toTest), "i")
	ifSuffixN := strings.HasSuffix(strings.ToLower(toTest), "n")
	ifContainsA := strings.Contains(strings.ToLower(toTest), "a")

	//Simple control flow conditional
	if ifPrefixI && ifSuffixN && ifContainsA {
		fmt.Print("Found!\n")
	} else {
		fmt.Print("Not Found!\n")
	}

}
