package main

import (
	"fmt"
	"github.com/jteeuwen/glfw"
)


type Window struct {
	
}

func (w *Window) openWindow() bool {
	var err error
	
	// init
	err = glfw.Init()
	if err != nil {
		fmt.Printf("glfw.Init: %s", err)
		return false
	}
	
	// windowed or fullscreen?
	windowMode := glfw.Windowed
	if *flagFullscreen {
		windowMode = glfw.Fullscreen
	}
	
	// open window
	err = glfw.OpenWindow(*flagWidth, *flagHeight, 8, 8, 8, 8, 24, 0, windowMode)
	if err != nil {
		fmt.Printf("glfw.OpenWindow: %s", err)
		return false
	}
	
	// Ensure window is cleanly closed on exit.
	defer glfw.CloseWindow()

	glfw.SetWindowTitle("Minecraft Go")
	
	// set window position if defined
	if *flagXpos > -1 && *flagYpos > -1 {
		glfw.SetWindowPos(*flagXpos, *flagYpos)
	}
	
	// callbacks
	glfw.SetWindowSizeCallback(onResize)
	glfw.SetWindowCloseCallback(onClose)

	return true
}

func onResize(width, height int) {
	
}

func onClose() int {
	Close()
	return 1
}

