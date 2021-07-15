package main

import (
	"fmt"
)

type Point struct {
	x int32
	y int32
}

func main() {
	var p1 Point = Point{0, 0}
	var p2 Point = Point{2, 2}

	fmt.Println("p1.x=", p1.x, "p1.y=", p1.y)
	fmt.Println("p2.x=", p2.x, "p2.y=", p2.y)

	p1_memory := &p1
	p1_memory.x = 33
	// p1_memory.y = 44
	p1.x = 45
	p1.y = 33

	fmt.Println("p1.x=", p1.x, "p1.y=", p1.y)
	fmt.Println("p2.x=", p2.x, "p2.y=", p2.y)

	// fmt.Println("p1_mem", p1_memory.x )
}
