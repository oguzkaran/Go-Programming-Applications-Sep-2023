/*
------------------------------------------------------------------------------------------------------------------------




------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

type Point struct {
	x, y int
}

func (p *Point) Set(x, y int) {
	p.x = x
	p.y = y
}

func (p *Point) Offset(dx, dy int) {
	p.x += dx
	p.y += dy
}

func (p *Point) Print() {
	fmt.Printf("x = %d, y = %d\n", p.x, p.y)
}

func New(x, y int) *Point {
	return &Point{x, y}
}

func main() {
	var x, y int
	fmt.Print("Input coordinates:")
	fmt.Scan(&x, &y)

	p1 := New(x, y)
	p2 := *p1
	p1.Print()
	p2.Print()
	p1.Offset(20, -30)
	p1.Print()
	p2.Print()
}
