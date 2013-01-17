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
	switch key {
	case glfw.KeyEsc:
		input.window.Stop()
	case glfw.KeyUp:
		if state == glfw.KeyRelease {
			input.camera.LogPosition()
			
			input.camera.PosX += input.camera.DirX
			input.camera.PosY += input.camera.DirY
			input.camera.PosZ += input.camera.DirZ
			input.camera.NewPosition = true
			
			input.camera.LogPosition()
		}
	case glfw.KeyDown:
		if state == glfw.KeyRelease {
			input.camera.LogPosition()
			
			
			input.camera.NewPosition = true
			
			input.camera.LogPosition()
		}
	case glfw.KeyLeft:
		if state == glfw.KeyRelease {
			input.camera.LogPosition()
			
			
			input.camera.NewPosition = true
			
			input.camera.LogPosition()
		}
	case glfw.KeyRight:
		if state == glfw.KeyRelease {
			input.camera.LogPosition()
			
			
			input.camera.NewPosition = true
			
			input.camera.LogPosition()
		}
	}
}

func onChar(key, state int) {
//	println("char", key, state)
}

