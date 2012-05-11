package main

import (
	"fmt"
	"github.com/jteeuwen/glfw"
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"flag"
	"time"
)

var width = flag.Int("width", 800, "width of the window")
var height = flag.Int("height", 600, "height of the window")
var fullscreen = flag.Bool("fullscreen", false, "enable fullscreen")

func main() {
	fmt.Println("starting")
	flag.Parse()

	openWindow()
	initGL()
	
	for {
		drawScene()
	}
	
	time.Sleep(5 * time.Second)
	
	glfw.CloseWindow()
	glfw.Terminate()
	
	fmt.Println("the end")
}

func openWindow() bool {
	var err error
	
	err = glfw.Init()
	if err != nil {
		fmt.Printf("glfw.Init: %s", err)
		return false
	}
	
	windowMode := glfw.Windowed
	if *fullscreen {
		windowMode = glfw.Fullscreen
		(*width) = 0
		(*height) = 0
	}
	err = glfw.OpenWindow(*width, *height, 8, 8, 8, 8, 24, 0, windowMode)
	if err != nil {
		fmt.Printf("glfw.OpenWindow: %s", err)
		return false
	}

	glfw.SetWindowTitle("Minecraft Go")

	return true
}

func initGL() {
	glfw.SetSwapInterval(1)
	
	gl.ShadeModel(gl.SMOOTH)
	gl.ClearColor(0, 0, 0, 0)
	gl.ClearDepth(1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)
	
	gl.Viewport(0, 0, *width, *height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(45.0, float64(*width)/float64(*height), 0.1, 100.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()
	
	gl.Translatef(0, 0, 0)
	t := time.Now()
	var millisecond float32 = (float32)(t.Second() * 1000 + t.Nanosecond() / 1000000)
//	gl.Rotatef((float32)(((float32(time.Now().Nanosecond()) / 1000000000.0) * 360)), 0, 1, 0)
	gl.Rotatef(10, 1, 0, 0)
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

func waitEvents(keepRunning chan bool) {
	for <- keepRunning {
		
	}
}
