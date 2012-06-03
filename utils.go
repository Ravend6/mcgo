package main

import (
	"github.com/jteeuwen/glfw"
)


// stores a position
type Vector struct {
	X, Y, Z float64
}


// returns milliseconds
func GetTime() int {
	return int(glfw.Time() * 1000)
}


