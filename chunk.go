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
		if (*posX) == 255 {
			// TODO do something
		} else {
			if c.blocks[(*posX)+1][(*posY)][(*posZ)] != nil {
				c.blocks[(*posX)+1][(*posY)][(*posZ)].SetVisibleSideMinusX(true)
			}
		}
		
		// check for lower edge
		if (*posX) == 0 {
			// TODO do something	
		} else {
			if c.blocks[(*posX)][(*posY)-1][(*posZ)] != nil {
				c.blocks[(*posX)][(*posY)-1][(*posZ)].SetVisibleSidePlusY(true)
			}
		}
		// check for higher edge
		if (*posX) == 255 {
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
		if (*posX) == 255 {
			// TODO do something
		} else {
			if c.blocks[(*posX)][(*posY)][(*posZ)+1] != nil {
				c.blocks[(*posX)][(*posY)][(*posZ)+1].SetVisibleSideMinusZ(true)
			}
		}
	} else {
		
	}
}
