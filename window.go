package main

import (
	"fmt"
	"github.com/jteeuwen/glfw"
	"flag"
	"time"
)

var width = flag.Int("width", 800, "width of the window")
var height = flag.Int("height", 600, "height of the window")
var fullscreen = flag.Bool("fullscreen", false, "enable fullscreen")

func main() {
	fmt.Println("starting")
	flag.Parse()
	
	var err error
	
	err = glfw.Init()
	if err != nil {
		fmt.Printf("glfw.Init: %s", err)
		return;
	}
	
	windowMode := glfw.Windowed
	if *fullscreen {
		windowMode = glfw.Fullscreen
	}
	err = glfw.OpenWindow(*width, *height, 8, 8, 8, 8, 0, 0, windowMode)
	if err != nil {
		fmt.Printf("glfw.OpenWindow: %s", err)
		return;
	}
	
	time.Sleep(5 * time.Second)
	
	glfw.CloseWindow()
	glfw.Terminate()
	
	fmt.Println("the end")
}

