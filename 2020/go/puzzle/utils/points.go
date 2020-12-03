package utils

import "math"

// ManhattanDistance computes the Manhattan distance between 2 points
func ManhattanDistance(p1 Point2D, p2 Point2D) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}

// Point2D represents a point in 2 dimensions with integer coordinates
type Point2D struct {
	X int
	Y int
}

func (p *Point2D) Add(another Point2D) *Point2D {
	p.X += another.X
	p.Y += another.Y
	return p
}

// ManhattanDistance computes the manhattan distance between this point and another
func (p Point2D) ManhattanDistance(otherPoint Point2D) int {
	return ManhattanDistance(p, otherPoint)
}

// Vector2D is a vector in 2 dimensions
type Vector2D struct {
	Origin      Point2D
	Destination Point2D
}

func (v Vector2D) ComponentX() int {
	return v.Destination.X - v.Origin.X
}

func (v Vector2D) ComponentY() int {
	return v.Destination.Y - v.Origin.Y
}

func (v Vector2D) Norm() float64 {
	dx, dy := v.ComponentX(), v.ComponentY()
	return math.Sqrt(math.Pow(float64(dx), 2) + math.Pow(float64(dy), 2))
}

func (v Vector2D) DotProduct(u Vector2D) int {
	dx1, dy1 := v.ComponentX(), v.ComponentY()
	dx2, dy2 := u.ComponentX(), u.ComponentY()
	return dx1*dx2 + dy1*dy2
}

func (v Vector2D) AngleWith(u Vector2D) float64 {
	dotProd := v.DotProduct(u)
	n1, n2 := v.Norm(), u.Norm()
	return math.Acos(float64(dotProd) / float64(n1*n2))
}
