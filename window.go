package main

import (
	"github.com/jteeuwen/glfw"
	"time"
)


type Window struct {
	renderer Renderer
	world World
	input Input
	camera Camera
	
	keepRunning bool
}

var window *Window

func (w *Window) Start() {
	window = w
	
	println("Window::Start")
	defer println("END Window::Start")
	
	defer w.Close()
	
	if !w.openWindow() {
		return
	}
	
	// callback
	glfw.SetWindowCloseCallback(onClose)
	
	w.world.Init(&w.renderer)
	
	w.renderer.Init(&w.world, &w.camera)
	w.renderer.InitGL()
	
	w.input.Init(w)
	
	// run threads
	go w.renderer.Start()
	go w.input.Start()
	
	w.keepRunning = true
	
	for w.keepRunning {
		time.Sleep(100 * time.Millisecond)
	}
}

func (w *Window) Close() {
	w.Stop()
	glfw.CloseWindow()
	glfw.Terminate()
}

func (w *Window) Stop() {
	w.keepRunning = false
	w.renderer.Stop()
	time.Sleep(500 * time.Millisecond)
}

func (w *Window) openWindow() bool {
	var err error
	
	// init
	err = glfw.Init()
	if err != nil {
		print("glfw.Init: %s\n", err)
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
		print("glfw.OpenWindow: %s\n", err)
		return false
	}
	
	glfw.SetWindowTitle("Minecraft Go")
	
	// set window position if defined
	if *flagXpos > -1 && *flagYpos > -1 {
		glfw.SetWindowPos(*flagXpos, *flagYpos)
	}
	
	return true
}

func onClose() int {
	window.Stop()
	return 1
}

