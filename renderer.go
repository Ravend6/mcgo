package main

import (
//	"fmt"
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"github.com/jteeuwen/glfw"
	"time"
)


func initGL() {
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
		0, 0, 0, // position
		0, 0, 1, // direction
		0, 1, 0) // up
	
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()
	
	gl.Translatef(0, 0, 4)
	t := time.Now()
	millisecond := (float32)(t.Second() * 1000 + t.Nanosecond() / 1000000)
//	gl.Rotatef(10, 1, 0, 0)
	gl.Rotatef((millisecond / 1000) * 360, 0, 1, 0)
	
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
	
	glfw.SwapBuffers()
}

