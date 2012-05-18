package main

import (
	"fmt"
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
	
	var mWindow Window

	if !mWindow.openWindow() {
		return;
	}
	
	var mRenderer Renderer
	
	mRenderer.initGL()
	
	keepRunning := true
	
	// TODO put this in window?
	for keepRunning {
		mRenderer.drawScene()
		
		// TODO put this in window and input handler
		if glfw.Key(glfw.KeyEsc) == glfw.KeyPress || glfw.WindowParam(glfw.Opened) == 0 {
			keepRunning = false
		}
	}
	
	glfw.CloseWindow()
	glfw.Terminate()
	
	fmt.Println("the end")
}

