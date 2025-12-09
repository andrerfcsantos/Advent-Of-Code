package geometry

type Vector2D struct {
	DX int
	DY int
}

func (v *Vector2D) ComponentX() int {
	return v.DX
}

func (v *Vector2D) ComponentY() int {
	return v.DY
}

type Vector3D struct {
	DX int
	DY int
	DZ int
}
