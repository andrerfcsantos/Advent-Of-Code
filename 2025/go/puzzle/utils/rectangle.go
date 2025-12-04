package utils

// Rectangle represents a rectangle in a grid of integer coordinates
type Rectangle struct {
	Corner1 Point2D
	Corner2 Point2D
}

// TopLeftCorner returns the top bottom corner of the rectangle
func (r *Rectangle) TopLeftCorner() Point2D {
	return Point2D{
		X: Min(r.Corner1.X, r.Corner2.X),
		Y: Max(r.Corner1.Y, r.Corner2.Y),
	}
}

// TopRightCorner returns the top right corner of the rectangle
func (r *Rectangle) TopRightCorner() Point2D {
	return Point2D{
		X: Max(r.Corner1.X, r.Corner2.X),
		Y: Max(r.Corner1.Y, r.Corner2.Y),
	}
}

// BottomRightCorner returns the bottom right corner of the rectangle
func (r *Rectangle) BottomRightCorner() Point2D {
	return Point2D{
		X: Max(r.Corner1.X, r.Corner2.X),
		Y: Min(r.Corner1.Y, r.Corner2.Y),
	}
}

// BottomLeftCorner returns the bottom left corner of the rectangle
func (r *Rectangle) BottomLeftCorner() Point2D {
	return Point2D{
		X: Min(r.Corner1.X, r.Corner2.X),
		Y: Min(r.Corner1.Y, r.Corner2.Y),
	}
}

// AmplitudeX returns how large is the rectangle in the X (horizontal) axis
func (r *Rectangle) AmplitudeX() int {
	min, max := MinMax(r.Corner1.X, r.Corner2.X)
	return Abs(max - min)
}

// AmplitudeY returns how large is the rectangle in the Y (vertical) axis
func (r *Rectangle) AmplitudeY() int {
	min, max := MinMax(r.Corner1.Y, r.Corner2.Y)
	return Abs(max - min)
}
