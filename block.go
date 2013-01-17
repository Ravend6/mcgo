package main


type Block struct {
/*	xz uint8
	y uint8*/ // don't need them, get the position with position in chunk
	
//	data uint // size??
	
	visibleSides uint8 // +x -x +y -y +z -z
	
	parentChunk *Chunk
}

func (b *Block) Init(parentChunk *Chunk/*, x uint8, y uint8, z uint8*/) {
	b.parentChunk = parentChunk
	
/*	b.y = y
	b.xz = x + z*16*/
	
	b.visibleSides = 0
}
/*
func (b *Block) GetPosition() (x uint8, y uint8, z uint8) {
	return b.xz%16, b.y, b.xz/16
}*/

func (b *Block) SetVisibleSidePlusX(visible bool) {
	if visible {
		b.visibleSides |= uint8(0x1)
	} else {
		b.visibleSides &= ^uint8(0x1)
	}
}

func (b *Block) SetVisibleSideMinusX(visible bool) {
	if visible {
		b.visibleSides |= uint8(0x2)
	} else {
		b.visibleSides &= ^uint8(0x2)
	}
}

func (b *Block) SetVisibleSidePlusY(visible bool) {
	if visible {
		b.visibleSides |= uint8(0x4)
	} else {
		b.visibleSides &= ^uint8(0x4)
	}
}

func (b *Block) SetVisibleSideMinusY(visible bool) {
	if visible {
		b.visibleSides |= uint8(0x8)
	} else {
		b.visibleSides &= ^uint8(0x8)
	}
}

func (b *Block) SetVisibleSidePlusZ(visible bool) {
	if visible {
		b.visibleSides |= uint8(0x10)
	} else {
		b.visibleSides &= ^uint8(0x10)
	}
}

func (b *Block) SetVisibleSideMinusZ(visible bool) {
	if visible {
		b.visibleSides |= uint8(0x20)
	} else {
		b.visibleSides &= ^uint8(0x20)
	}
}

func (b *Block) IsVisible() (visible bool) {
	if b.GetVisibleSides() == 0 {
		return false
	}
	return true
}

func (b *Block) GetVisibleSidesBool() (plusX, minusX, plusY, minusY, plusZ, minusZ bool) {
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

func (b *Block) GetVisibleSides() uint8 {
	return (b.visibleSides & 0x3f)
}

