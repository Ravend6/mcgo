package main

import (
	"github.com/go-gl/glfw"
	"time"
)

type Input struct {
	camera *Camera
	window *Window
	
	keepRunning bool
}

var input *Input

func (i *Input) Init(window *Window, camera *Camera) {
	input = i
	i.window = window
	i.camera = camera
	
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

