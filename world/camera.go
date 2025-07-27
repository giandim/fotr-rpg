package world

type Camera struct {
	X      int
	Y      int
	width  int
	height int
}

func NewCamera(x, y, width, height int) *Camera {
	return &Camera{
		X:      x,
		Y:      y,
		width:  width,
		height: height,
	}
}

func (c *Camera) UpdatePosition(x int, y int) {
	c.X = x
	c.Y = y
}
