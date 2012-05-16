package main

import (
	"fmt"
	"github.com/jteeuwen/glfw"
	"flag"
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
	
	var mRenderer Renderer
	
	mRenderer.initGL()
	
	keepRunning := true
	
	for keepRunning {
		mRenderer.drawScene()
		
		if glfw.Key(glfw.KeyEsc) == glfw.KeyPress || glfw.WindowParam(glfw.Opened) == 0 {
			keepRunning = false
		}
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

func waitEvents(keepRunning chan bool) {
	for <- keepRunning {
		
	}
}
