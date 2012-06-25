package main

import (
	"container/list"
	"math/rand"
)

type World struct {
	Chunklist list.List
	
	renderer *Renderer
}

func (w *World) Init(rend *Renderer) {
	w.renderer = rend
	
	var foo Chunk
	chunk := &foo
	chunk.x = 0
	chunk.z = 0
	
	w.Chunklist.PushBack(chunk)
	
	count := 0
	
	for indexX, _ := range chunk.blocks {
		for indexY, _ := range chunk.blocks[indexX] {
			for indexZ, _ := range chunk.blocks[indexX][indexY] {
				if rand.Intn(4) == 0 {
					count++
					var b Block
					b.Init(chunk)
					chunk.blocks[indexX][indexY][indexZ] = &b
				}
			}
		}
	}
	
	chunk.UpdateAllVisibleSides()
	
	w.renderer.CreateChunk(chunk)
	
	println("Generated blocks in chunk:", count)
}


