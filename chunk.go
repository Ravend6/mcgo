package main

type Chunk struct {
	x, z int
	blocks [16][256][16]Block
	
	
}

func (c *Chunk) GetPosition() (x int, z int) {
	return c.x, c.z
}

func (c *Chunk) ToVertices() () {

}

