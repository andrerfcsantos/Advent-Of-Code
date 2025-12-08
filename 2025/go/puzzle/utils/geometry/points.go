package geometry

import "math"

type Point2D struct {
	X int
	Y int
}

func (p *Point2D) Add(vector Vector2D) Point2D {
	return Point2D{
		X: p.X + vector.DX,
		Y: p.Y + vector.DY,
	}
}

type Point3D struct {
	X int
	Y int
	Z int
}

func (p *Point3D) EuclideanDistanceTo(another Point3D) float64 {
	dx := float64(another.X - p.X)
	dy := float64(another.Y - p.Y)
	dz := float64(another.Z - p.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (p *Point3D) Add(vector Vector3D) Point3D {
	return Point3D{
		X: p.X + vector.DX,
		Y: p.Y + vector.DY,
		Z: p.Z + vector.DZ,
	}
}
