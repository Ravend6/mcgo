package main

import (
	"fmt"
	"github.com/jteeuwen/glfw"
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"flag"
	"time"
)

var flagWidth = flag.Int("width", 800, "width of the window")
var flagHeight = flag.Int("height", 600, "height of the window")
var flagXpos = flag.Int("xpos", -1, "default x position of the window (-1 for none)")
var flagYpos = flag.Int("ypos", -1, "default y position of the window (-1 for none)")
var flagFullscreen = flag.Bool("fullscreen", false, "enable fullscreen")

func main() {
	fmt.Println("starting")
	flag.Parse()

	if !openWindow() {
		return;
	}
	
	initGL()
	
	keepRunning := true
	
	for keepRunning {
		drawScene()
	}
	
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
	if *flagFullscreen {
		windowMode = glfw.Fullscreen
		(*flagWidth) = 0
		(*flagHeight) = 0
	}
	err = glfw.OpenWindow(*flagWidth, *flagHeight, 8, 8, 8, 8, 24, 0, windowMode)
	if err != nil {
		fmt.Printf("glfw.OpenWindow: %s", err)
		return false
	}

	glfw.SetWindowTitle("Minecraft Go")
	
	if *flagXpos > -1 && *flagYpos > -1 {
		glfw.SetWindowPos(*flagXpos, *flagYpos)
	}

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
	
	width, height := glfw.WindowSize()
	
	gl.Viewport(0, 0, width, height)
	
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(45.0, float64(width)/float64(height), 0.1, 100.0)
	
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()
	
	gl.Translatef(1, 0, -4)
	t := time.Now()
	millisecond := (float32)(t.Second() * 1000 + t.Nanosecond() / 1000000)
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
