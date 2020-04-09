package main

import "fmt"

func main() {
	b := initGameBoard()
	fmt.Println("Lets go!")
	fmt.Println(b.getAllEmptyLocations())
}
