package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Please input the file location:")

	//scannes for file address
	scanner := bufio.NewReader(os.Stdin)
	address, _ := scanner.ReadString('\n')

	//Opening the file
	file, err := os.Open(strings.TrimRight(address, "\n"))
	if err != nil {
		log.Fatalf("Error Occured : %s ", err)
	}

	//reader for file
	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)

	//slice of name struct
	fileData := make([]name, 0)

	//scanning through the lines one by one
	for reader.Scan() {
		arr := strings.Split(reader.Text(), " ")
		fileData = append(fileData, name{fname: arr[0], lname: arr[1]})
	}

	//important to close the file
	file.Close()

	//Printing the data stored in slice
	for _, line := range fileData {
		fmt.Println("fname: ", line.fname, ", lname: ", line.lname)
	}
}
