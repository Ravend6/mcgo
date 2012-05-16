package main

type Chunk struct {
	x, z int
	blocks [16][256][16]*Block
	
	
}

func (c *Chunk) GetPosition() (x int, z int) {
	return c.x, c.z
}

func (c *Chunk) GetBlocks() (blocks *[16][256][16]*Block) {
	return &c.blocks
}

func (c *Chunk) CheckVertices() {
	
	for indexX, _ := range c.blocks {
		for indexY, _ := range c.blocks[indexX] {
			for indexZ, value := range c.blocks[indexX][indexY] {
				
			}
		}
	}
}

