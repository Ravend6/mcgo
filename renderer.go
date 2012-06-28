package main

import (
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"github.com/jteeuwen/glfw"
)


type Renderer struct {
	timeCurrent, timeNextSecond, frames int
	
	world *World
	camera *Camera
	
	NewWindowSize bool
	
	keepRunning bool
}

var renderer *Renderer


func (r *Renderer) Init(world *World, camera *Camera) {
	renderer = r
	
	r.world = world
	r.camera = camera
	
	r.NewWindowSize = false
	
	r.keepRunning = true
	
	// callback
	glfw.SetWindowSizeCallback(onResize)
}

func (r *Renderer) InitGL() {
	if *flagVSync {
		glfw.SetSwapInterval(1)
	}
	
	gl.ShadeModel(gl.SMOOTH)
	gl.ClearColor(0, 0, 0, 0)
	gl.ClearDepth(1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)
	
	width, height := glfw.WindowSize()
	
	if height == 0 {
		height = 1
	}
	
	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(45.0, float64(width)/float64(height), 0.1, 100.0)
	glu.LookAt(
		r.camera.PosX, r.camera.PosY, r.camera.PosZ, // position
		r.camera.DirX, r.camera.DirY, r.camera.DirZ, // direction
		r.camera.UpX,  r.camera.UpY,  r.camera.UpZ ) // up
	
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	
	r.timeNextSecond = 1000
	r.frames = 0
}

func (r *Renderer) createBlock(block *Block) {
	if block == nil || block.IsVisible() == false {
		return
	}
	
	plusX, minusX, plusY, minusY, plusZ, minusZ := block.GetVisibleSides()
	
	block.DisplayList = gl.GenLists(1)
	gl.NewList(block.DisplayList, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, []float32{0.5, 0.5, 0.5, 1.0})
	
	// +x
	if plusX {
		gl.Color3f(1, 0, 0)
		gl.Begin(gl.QUADS)
		gl.Vertex3f(0.5, -0.5, 0.5)
		gl.Vertex3f(0.5, -0.5, -0.5)
		gl.Vertex3f(0.5, 0.5, -0.5)
		gl.Vertex3f(0.5, 0.5, 0.5)
		gl.End()
	}
	
	// -x
	if minusX {
		gl.Color3f(0, 1, 0)
		gl.Begin(gl.QUADS)
		gl.Vertex3f(-0.5, -0.5, -0.5)
		gl.Vertex3f(-0.5, 0.5, -0.5)
		gl.Vertex3f(-0.5, 0.5, 0.5)
		gl.Vertex3f(-0.5, -0.5, 0.5)
		gl.End()
	}
	
	// +y
	if plusY {
		gl.Color3f(0, 1, 1)
		gl.Begin(gl.QUADS)
		gl.Vertex3f(-0.5, -0.5, 0.5)
		gl.Vertex3f(-0.5, 0.5, 0.5)
		gl.Vertex3f(0.5, 0.5, 0.5)
		gl.Vertex3f(0.5, -0.5, 0.5)
		gl.End()
	}
	
	// -y (?)
	if minusY {
		gl.Color3f(0, 0, 1)
		gl.Begin(gl.QUADS)
		gl.Vertex3f(-0.5, -0.5, -0.5)
		gl.Vertex3f(-0.5, -0.5, 0.5)
		gl.Vertex3f(0.5, -0.5, 0.5)
		gl.Vertex3f(0.5, -0.5, -0.5)
		gl.End()
	}
	
	// +z (?)
	if plusZ {
		gl.Color3f(1, 0, 1)
		gl.Begin(gl.QUADS)
		gl.Vertex3f(-0.5, 0.5, -0.5)
		gl.Vertex3f(0.5, 0.5, -0.5)
		gl.Vertex3f(0.5, -0.5, -0.5)
		gl.Vertex3f(-0.5, -0.5, -0.5)
		gl.End()
	}
	
	// -z (?)
	if minusZ {
		gl.Color3f(1, 1, 0)
		gl.Begin(gl.QUADS)
		gl.Vertex3f(-0.5, 0.5, -0.5)
		gl.Vertex3f(-0.5, 0.5, 0.5)
		gl.Vertex3f(0.5, 0.5, 0.5)
		gl.Vertex3f(0.5, 0.5, -0.5)
		gl.End()
	}
	
	gl.EndList()
}

func (r *Renderer) CreateChunk(chunk *Chunk) {
	sw := Stopwatch()
	
	if chunk == nil {
		return
	}
	
	for indexX, _ := range chunk.blocks {
		for indexY, _ := range chunk.blocks[indexX] {
			for _, block := range chunk.blocks[indexX][indexY] {
				if block != nil {
					if block.IsVisible() {
						r.createBlock(block)
					}
				}
			}
		}
	}
	
	println("CreateChunk time:", sw.Stop(), "ms")
}

func (r *Renderer) drawScene() {
//	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
//	gl.LoadIdentity()
	
	glfw.SwapBuffers()
	
	r.frames++
	r.timeCurrent = GetTime()
	if r.timeCurrent >= r.timeNextSecond {
		println(r.frames, "fps")
		
		r.frames = 0
		r.timeNextSecond = r.timeNextSecond + 1000
	}
}

func (r *Renderer) Start() {
	println("Renderer::Start")
	defer println("END Renderer::Start")
	
	for r.keepRunning {
		if r.NewWindowSize == true {
			r.NewWindowSize = false
			
			width, height := glfw.WindowSize()
			
			if height == 0 {
				height = 1
			}
			
			gl.Viewport(0, 0, width, height)
/*			gl.MatrixMode(gl.PROJECTION)
			gl.LoadIdentity()*/
			glu.Perspective(45.0, float64(width)/float64(height), 0.1, 100.0)
		}
		
		if r.camera.NewPosition == true {
			r.camera.NewPosition = false
			glu.LookAt(
				r.camera.PosX, r.camera.PosY, r.camera.PosZ, // position
				r.camera.DirX, r.camera.DirY, r.camera.DirZ, // direction
				r.camera.UpX,  r.camera.UpY,  r.camera.UpZ ) // up
		}
		
		r.drawScene()
	}
}

func (r *Renderer) Stop() {
	r.keepRunning = false
}


  ///////////////
 // Callbacks //
///////////////

func onResize(width, height int) {
	renderer.NewWindowSize = true
}

