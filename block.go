package main


type Block struct {
	xz uint8
	y uint8
//	data uint // size??
	
	visibleSides uint8 // +x -x +y -y +z -z
	
	parentChunk *Chunk
	
}

func (b *Block) Init(parentChunk *Chunk, x uint8, y uint8, z uint8) {
	b.parentChunk = parentChunk
	
	b.y = y
	b.xz = x + z*16
	
	b.visibleSides = 255
}

func (b *Block) GetPosition() (x uint8, y uint8, z uint8) {
	return b.xz%16, b.y, b.xz/16
}

func (b *Block) SetVisibleSides(plusX bool, minusX bool, plusY bool, minusY bool, plusZ bool, minusZ bool) {
	b.visibleSides = 0
	if plusX {
		b.visibleSides |= 0x1
	}
	if minusX {
		b.visibleSides |= 0x2
	}
	if plusY {
		b.visibleSides |= 0x4
	}
	if minusY {
		b.visibleSides |= 0x8
	}
	if plusZ {
		b.visibleSides |= 0x10
	}
	if minusZ {
		b.visibleSides |= 0x20
	}
}

func (b *Block) GetVisibleSides() (plusX bool, minusX bool, plusY bool, minusY bool, plusZ bool, minusZ bool) {
	if (b.visibleSides & 0x1) > 0 {
		plusX = true
	}
	if (b.visibleSides & 0x2) > 0 {
		minusX = true
	}
	if (b.visibleSides & 0x4) > 0 {
		plusY = true
	}
	if (b.visibleSides & 0x8) > 0 {
		minusY = true
	}
	if (b.visibleSides & 0x10) > 0 {
		plusZ = true
	}
	if (b.visibleSides & 0x20) > 0 {
		minusZ = true
	}
	
	return
}

