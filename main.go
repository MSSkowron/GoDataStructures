package main

import (
	"fmt"

	"github.com/MSSkowron/GoDataStructuresAndAlgorithms/stack"
)

func main() {
	s := stack.New[int]()

	fmt.Printf("%+v\n", s)
}
