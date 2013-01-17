package main

import (
)


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

func (c *Chunk) UpdateAllVisibleSides() {
	sw := Stopwatch()
	
	var zStartSecond bool = false
	// go through all block
	for indexX, _ := range c.blocks {
		for indexY, _ := range c.blocks[indexX] {
			var indexZ int
			// select every other block
			if zStartSecond {
				indexZ = 1
				zStartSecond = false
			} else {
				indexZ = 0
				zStartSecond = true
			}
			for ; indexZ < 16; indexZ = indexZ+2 {
				c.UpdateBlockVisibleSides(&indexX, &indexY, &indexZ)
			}
		}
	}
	println("UpdateAllVisibleSides time:", sw.Stop(), "ms")
	
	var numSides = 0
	for indexX, _ := range c.blocks {
		for indexY, _ := range c.blocks[indexX] {
			for indexZ, _ := range c.blocks[indexX][indexY] {
				if c.blocks[indexX][indexY][indexZ] != nil {
					plusX, minusX, plusY, minusY, plusZ, minusZ := c.blocks[indexX][indexY][indexZ].GetVisibleSidesBool()
					if plusX {
						numSides++
					}
					if minusX {
						numSides++
					}
					if plusY {
						numSides++
					}
					if minusY {
						numSides++
					}
					if plusZ {
						numSides++
					}
					if minusZ {
						numSides++
					}
				}
			}
		}
	}
	println("num visible sides: ", numSides)
}

func (c *Chunk) UpdateBlockVisibleSides(posX, posY, posZ *int) {
	// check if block exists
	if c.blocks[(*posX)][(*posY)][(*posZ)] == nil {
		// tell all surrounding squares to be visible
		
		// check for lower edge
		if (*posX) == 0 {
			// TODO do something
		} else {
			if c.blocks[(*posX)-1][(*posY)][(*posZ)] != nil {
				c.blocks[(*posX)-1][(*posY)][(*posZ)].SetVisibleSidePlusX(true)
			}
		}
		// check for higher edge
		if (*posX) == 15 {
			// TODO do something
		} else {
			if c.blocks[(*posX)+1][(*posY)][(*posZ)] != nil {
				c.blocks[(*posX)+1][(*posY)][(*posZ)].SetVisibleSideMinusX(true)
			}
		}
		
		// check for lower edge
		if (*posY) == 0 {
			// TODO do something
		} else {
			if c.blocks[(*posX)][(*posY)-1][(*posZ)] != nil {
				c.blocks[(*posX)][(*posY)-1][(*posZ)].SetVisibleSidePlusY(true)
			}
		}
		// check for higher edge
		if (*posY) == 255 {
			// TODO do something
		} else {
			if c.blocks[(*posX)][(*posY)+1][(*posZ)] != nil {
				c.blocks[(*posX)][(*posY)+1][(*posZ)].SetVisibleSideMinusY(true)
			}
		}
		
		// check for lower edge
		if (*posZ) == 0 {
			// TODO do something
		} else {
			if c.blocks[(*posX)][(*posY)][(*posZ)-1] != nil {
				c.blocks[(*posX)][(*posY)][(*posZ)-1].SetVisibleSidePlusZ(true)
			}
		}
		// check for higher edge
		if (*posZ) == 15 {
			// TODO do something
		} else {
			if c.blocks[(*posX)][(*posY)][(*posZ)+1] != nil {
				c.blocks[(*posX)][(*posY)][(*posZ)+1].SetVisibleSideMinusZ(true)
			}
		}
	} else { // block exists
		// check for -x edge
		if (*posX) == 0 {
			// TODO do something
		} else {
			if c.blocks[(*posX)-1][(*posY)][(*posZ)] == nil {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSideMinusX(true)
			} else {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSideMinusX(false)
				c.blocks[(*posX)-1][(*posY)][(*posZ)].SetVisibleSideMinusX(false)
			}
		}
		
		// check for +x edge
		if (*posX) == 15 {
			// TODO do something
		} else {
			if c.blocks[(*posX)+1][(*posY)][(*posZ)] == nil {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSidePlusX(true)
			} else {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSidePlusX(false)
				c.blocks[(*posX)+1][(*posY)][(*posZ)].SetVisibleSidePlusX(false)
			}
		}
		
		// check for -y edge
		if (*posY) == 0 {
			c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSideMinusY(true)
		} else {
			if c.blocks[(*posX)][(*posY)-1][(*posZ)] == nil {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSideMinusY(true)
			} else {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSideMinusY(false)
				c.blocks[(*posX)][(*posY)-1][(*posZ)].SetVisibleSideMinusY(false)
			}
		}
		// check for +y edge
		if (*posY) == 255 {
			c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSidePlusY(true)
		} else {
			if c.blocks[(*posX)][(*posY)+1][(*posZ)] == nil {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSidePlusY(true)
			} else {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSidePlusY(false)
				c.blocks[(*posX)][(*posY)+1][(*posZ)].SetVisibleSidePlusY(false)
			}
		}
		
		// check for -z edge
		if (*posZ) == 0 {
			// TODO do something
		} else {
			if c.blocks[(*posX)][(*posY)][(*posZ)-1] == nil {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSideMinusZ(true)
			} else {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSideMinusZ(false)
				c.blocks[(*posX)][(*posY)][(*posZ)-1].SetVisibleSideMinusZ(false)
			}
		}
		// check for +z edge
		if (*posZ) == 15 {
			// TODO do something
		} else {
			if c.blocks[(*posX)][(*posY)][(*posZ)+1] == nil {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSidePlusZ(true)
			} else {
				c.blocks[(*posX)][(*posY)][(*posZ)].SetVisibleSidePlusZ(false)
				c.blocks[(*posX)][(*posY)][(*posZ)+1].SetVisibleSidePlusZ(false)
			}
		}
	}
}

