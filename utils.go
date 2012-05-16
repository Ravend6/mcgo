package main

import (
	"time"
)


// stores a position
type Vector struct {
	X, Y, Z float64
}


// returns milliseconds
func GetTime() float32 {
	t := time.Now()
	millisecond := (float32)(t.Second() * 1000 + t.Nanosecond() / 1000000)
	return millisecond
}

