package main


type Block struct {
	xz int8
	y int8
	data int // size??
	
	visible
	
}

func (b *Block) GetPosition() (x int8, y int8, z int8) {
	return b.xz%16, b.y, b.xz/16
}


