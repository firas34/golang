package main

import (
	"fmt"
	"math"
)

func longestOnes(nums []int, k int) int {
	return 5
}

type MyType struct {
	name   string
	number int
}

func (m *MyType) updateMyType(name string, num int) {
	m.name = name
	m.number = num
}

//interface

type Shape interface {
	area() float64
}

type Square struct {
	x1, y1, x2, y2 float64
}

func (s *Square) area() float64 {
	return math.Abs(s.x2-s.x1) * math.Abs(s.y2-s.y1)
}

func display(a int) {
	fmt.Println(a)
}

func main() {

	// var a MyType
	// a.updateMyType("Firas", 93181680)
	// s := Square{0, 0, 5, 5}

	// fmt.Println(s.area())
	go display(15)
	var in string
	fmt.Scanln(&in)

	fmt.Println(in)
}
