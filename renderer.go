package main

import (
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"github.com/jteeuwen/glfw"
)


type Renderer struct {
	timeCurrent, timeNextSecond, frames int
}

func (r *Renderer) initGL() {
	glfw.SetSwapInterval(1)
	
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
		0, 0, -50, // position
		0, 0, 1, // direction
		0, 1, 0) // up
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	
	r.timeNextSecond = 1000
	r.frames = 0
}

func (r *Renderer) drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()
	
	for i := 0; i < 10; i++ {
		gl.Translatef(1.2, 0, 0)
		gl.Rotatef((float32(GetTime()) / 1000.0) * 360, 0, 1, 0)
	
		gl.Color3f(1, 0, 1)
		gl.Begin(gl.QUADS) // a
		gl.Vertex3f(-0.5, 0.5, -0.5)
		gl.Vertex3f(0.5, 0.5, -0.5)
		gl.Vertex3f(0.5, -0.5, -0.5)
		gl.Vertex3f(-0.5, -0.5, -0.5)
		gl.End()
		gl.Color3f(0, 1, 0)
		gl.Begin(gl.QUADS) // b
		gl.Vertex3f(-0.5, -0.5, -0.5)
		gl.Vertex3f(-0.5, 0.5, -0.5)
		gl.Vertex3f(-0.5, 0.5, 0.5)
		gl.Vertex3f(-0.5, -0.5, 0.5)
		gl.End()
		gl.Color3f(0, 1, 1)
		gl.Begin(gl.QUADS) // c
		gl.Vertex3f(-0.5, -0.5, 0.5)
		gl.Vertex3f(-0.5, 0.5, 0.5)
		gl.Vertex3f(0.5, 0.5, 0.5)
		gl.Vertex3f(0.5, -0.5, 0.5)
		gl.End()
		gl.Color3f(1, 0, 0)
		gl.Begin(gl.QUADS) // d
		gl.Vertex3f(0.5, -0.5, 0.5)
		gl.Vertex3f(0.5, -0.5, -0.5)
		gl.Vertex3f(0.5, 0.5, -0.5)
		gl.Vertex3f(0.5, 0.5, 0.5)
		gl.End()
		gl.Color3f(0, 0, 1)
		gl.Begin(gl.QUADS) // e
		gl.Vertex3f(-0.5, -0.5, -0.5)
		gl.Vertex3f(-0.5, -0.5, 0.5)
		gl.Vertex3f(0.5, -0.5, 0.5)
		gl.Vertex3f(0.5, -0.5, -0.5)
		gl.End()
		gl.Color3f(1, 1, 0)
		gl.Begin(gl.QUADS) // f
		gl.Vertex3f(-0.5, 0.5, -0.5)
		gl.Vertex3f(-0.5, 0.5, 0.5)
		gl.Vertex3f(0.5, 0.5, 0.5)
		gl.Vertex3f(0.5, 0.5, -0.5)
		gl.End()
	}
	
	glfw.SwapBuffers()
	
	r.frames++
	r.timeCurrent = GetTime()
	if r.timeCurrent >= r.timeNextSecond {
		println(r.frames)
		
		r.frames = 0
		r.timeNextSecond = r.timeNextSecond + 1000
	}
}

