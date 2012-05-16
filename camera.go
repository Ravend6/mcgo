package main


type Camera struct {
	pos, dir, up Vector
}

func (c *Camera) Init(pos Vector, dir Vector, up Vector) {
	c.pos = pos
	c.dir = dir
	c.up = up
}

