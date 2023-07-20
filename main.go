package main

import (
	"fmt"

	"github.com/MSSkowron/GoDataStructesAndAlgorithms/stack"
)

func main() {
	s := stack.New[int]()

	fmt.Printf("%+v\n", s)
}
