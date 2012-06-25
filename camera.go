package main


type Camera struct {
	PosX, PosY, PosZ float64
	DirX, DirY, DirZ float64
	UpX, UpY, UpZ float64
	
	NewPosition bool
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
	
	c.NewPosition = true
}

func (c *Camera) SetPos(posX, posY, posZ float64) {
	c.PosX = posX
	c.PosY = posY
	c.PosZ = posZ
}

func (c *Camera) SetDir(dirX, dirY, dirZ float64) {
	c.DirX = dirX
	c.DirY = dirY
	c.DirZ = dirZ
}

func (c *Camera) SetUp(upX, upY, upZ float64) {
	c.UpX = upX
	c.UpY = upY
	c.UpZ = upZ
}

