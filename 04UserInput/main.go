package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// taking input

	// var name string
	print("Enter Your Name : ")

	// fmt.Scan(&name) // it does not read after whitespaces

	
	
	
	// to read even after whitespaces..
	
	reader := bufio.NewReader(os.Stdin)	
	name , _ := reader.ReadString('\n')  // it includes \n and prints it too
	fmt.Print("Hello, ", name)
}
