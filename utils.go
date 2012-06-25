package main

import (
	"github.com/jteeuwen/glfw"
)


// returns milliseconds
func GetTime() int {
	return int(glfw.Time() * 1000)
}


