package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//Defined a reader Object which reads from standard input
	reader := bufio.NewReader(os.Stdin)
	//Reading the name
	fmt.Println("Please input your name: ")
	name, _ := reader.ReadString('\n')
	//Reading the Address
	fmt.Println("Please input your address: ")
	address, _ := reader.ReadString('\n')

	//Making a dictionary
	var dict = map[string]string{"name": name, "address": address}

	//Making a []byte array
	jsonFormat, _ := json.Marshal(dict)

	fmt.Printf("\nJSON formatted data : \n")
	//Printing the JSON formatted dict
	fmt.Println(jsonFormat)

}
