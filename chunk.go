package main

import (
//	"fmt"
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

func (c *Chunk) CheckVertices() {
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
				// check if block exists
				if c.blocks[indexX][indexY][indexZ] == nil {
					// tell all surrounding squares to be visible
					
					// check for lower edge
					if indexX == 0 {
						// TODO do something	
					} else {
						if c.blocks[indexX-1][indexY][indexZ] != nil {
							c.blocks[indexX-1][indexY][indexZ].SetVisibleSidePlusX(true)
						}
					}
					// check for higher edge
					if indexX == 255 {
						// TODO do something
					} else {
						if c.blocks[indexX+1][indexY][indexZ] != nil {
							c.blocks[indexX+1][indexY][indexZ].SetVisibleSideMinusX(true)
						}
					}
					
					// check for lower edge
					if indexX == 0 {
						// TODO do something	
					} else {
						if c.blocks[indexX][indexY-1][indexZ] != nil {
							c.blocks[indexX][indexY-1][indexZ].SetVisibleSidePlusY(true)
						}
					}
					// check for higher edge
					if indexX == 255 {
						// TODO do something
					} else {
						if c.blocks[indexX][indexY+1][indexZ] != nil {
							c.blocks[indexX][indexY+1][indexZ].SetVisibleSideMinusY(true)
						}
					}
					
					// check for lower edge
					if indexZ == 0 {
						// TODO do something	
					} else {
						if c.blocks[indexX][indexY][indexZ-1] != nil {
							c.blocks[indexX][indexY][indexZ-1].SetVisibleSidePlusZ(true)
						}
					}
					// check for higher edge
					if indexX == 255 {
						// TODO do something
					} else {
						if c.blocks[indexX][indexY][indexZ+1] != nil {
							c.blocks[indexX][indexY][indexZ+1].SetVisibleSideMinusZ(true)
						}
					}
				} else {
					
				}
			}
		}
	}
}

