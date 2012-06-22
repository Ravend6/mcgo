package main

import (
	"github.com/jteeuwen/glfw"
)


type Window struct {
	mRenderer Renderer
	mWorld World
	mKeepRunning bool
}

func (w *Window) Start() {
	defer w.Close()
	
	if !w.openWindow() {
		return
	}
	
	w.mRenderer.Init(&w.mWorld)
	w.mRenderer.InitGL()
	
	w.mKeepRunning = true
	
	for w.mKeepRunning {
		w.mRenderer.drawScene()
	}
}

func (w *Window) Close() {
	w.mKeepRunning = false
	glfw.CloseWindow()
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
	
	// callbacks
/*	glfw.SetWindowSizeCallback(onResize)
	glfw.SetWindowCloseCallback(onClose)*/
	
	return true
}

func (w *Window) onResize(width, height int) {
	
}

func (w *Window) onClose() int {
	w.Close()
	return 1
}

