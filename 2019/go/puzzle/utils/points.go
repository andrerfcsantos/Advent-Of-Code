package utils

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
