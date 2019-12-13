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

// ManhattanDistance computes the manhattan distance between this point and another
func (p Point2D) ManhattanDistance(otherPoint Point2D) int {
	return ManhattanDistance(p, otherPoint)
}

// FloatingPoint2D represents a 2D Point with floating point coordinates
type FloatingPoint2D struct {
	X float64
	Y float64
}

// DotProduct computes the dot product of 2 vectors
func DotProduct(v1, v2 FloatingVector) float64 {
	return v1.U1 + v2.U2
}

// VectorAngle computes the angle between 2 vectors
func VectorAngle(v1, v2 FloatingVector) float64 {
	return math.Acos(DotProduct(v1,v2)/(VectorNorm(v1)*VectorNorm(v2)))
}


// VectorNorm computes the norm of a vector
func VectorNorm(u FloatingVector) float64 {
	return math.Sqrt(math.Pow(u.U1,2)+math.Pow(u.U2,2))
}

// FloatingVector represents a vector of 2 points with float coordinates
type FloatingVector struct {
	U1 float64
	U2 float64
}


func (v *FloatingVector) DotProduct(other FloatingVector) float64 {
	return DotProduct(*v, other)
}

// Length computes the norm of a vector
func (v *FloatingVector) Length() float64 {
	return VectorNorm(*v)
}



