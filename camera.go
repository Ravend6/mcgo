package main


type Camera struct {
	PosX, PosY, PosZ float64
	DirX, DirY, DirZ float64
	UpX, UpY, UpZ float64
	
	NewPosition int
}

func (c *Camera) Init(posX, posY, posZ, dirX, dirY, dirZ, upX, upY, upZ float64) {
	c.PosX = posX
	c.PosY = posY
	c.PosZ = posZ
	
	c.DirX = dirX
	c.DirY = dirY
	c.DirZ = dirZ
	
	c.UpX = upX
	c.UpY = upY
	c.UpZ = upZ
	
	c.NewPosition = 3
}

func (c *Camera) LogPosition() {
	println("Pos (", int(c.PosX), "/", int(c.PosY), "/", int(c.PosZ),
		") Dir (", int(c.DirX), "/", int(c.DirY), "/", int(c.DirZ),
		") Up (", int(c.UpX), "/", int(c.UpY), "/", int(c.UpZ), ")")
}

