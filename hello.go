package main

import "fmt"

func main() {
	var name string = "Dalvan"
	var age = 25
	version := 1.3
	fmt.Println("Hello Sr.", name, "Your age is", age)
	fmt.Println("This is the", version, "version")

	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit program")

	var command int
	fmt.Scan(&command)

	fmt.Println("The chosen command was", command)
}

//Posso usar := ao invés de decalrar var, como em name
