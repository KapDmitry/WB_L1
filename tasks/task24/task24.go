package main

import (
	"fmt"
	"math"
)

/*
Point - это структура для представления точек в 2D с параметрами x и y.
NewPoint - это конструктор для создания новой точки.
Distance - это метод структуры Point, который вычисляет расстояние между текущей точкой и переданной точкой в качестве аргумента
В main создаются две точки, вычисляется расстояние между ними, и результат выводится на экран.
*/
type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Distance(q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	distance := point1.Distance(point2)

	fmt.Printf("Point 1: (%.2f, %.2f)\n", point1.x, point1.y)
	fmt.Printf("Point 2: (%.2f, %.2f)\n", point2.x, point2.y)
	fmt.Printf("Distance between points: %.2f\n", distance)
}
