package main

import (
	"fmt"
)

type square struct {
	side int
}

func (s square) area() int {
	return s.side * s.side
}

type shape interface {
	area() int
}

func info(s shape){
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	newSquare := square{side: 20}
	info(newSquare)
}
