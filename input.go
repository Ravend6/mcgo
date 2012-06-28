package main

import (
	"github.com/jteeuwen/glfw"
	"time"
)

type Input struct {
	camera Camera
	window *Window
	
	keepRunning bool
}

var input *Input

func (i *Input) Init(window *Window) {
	input = i
	i.window = window
	
	i.camera.Init(
		0, 0, -50, // position
		0, 0, 1, // direction
		0, 1, 0) // up
	
	i.keepRunning = true
}

func (i *Input) Start() {
	// callbacks
	glfw.SetMouseButtonCallback(onMouseButton)
	glfw.SetMouseWheelCallback(onMouseWheel)
	glfw.SetKeyCallback(onKey)
	glfw.SetCharCallback(onChar)
	
	for i.keepRunning {
		time.Sleep(10 * time.Millisecond)
	}
}

func (i *Input) Stop() {
	i.keepRunning = false
}


  ///////////////
 // Callbacks //
///////////////


func onMouseButton(button, state int) {
	
}

func onMouseWheel(delta int) {
	
}

func onKey(key, state int) {
//	println("key", key, state)
	if key == glfw.KeyEsc {
		input.window.Stop()
	}
}

func onChar(key, state int) {
//	println("char", key, state)
}

