package main

import (
	"fmt"
	"fakedep"
)

func main() {
	fmt.Println("Hello, World from main")
	fakedep.PrintHelloWorld()
	fmt.Println("done!")
}
