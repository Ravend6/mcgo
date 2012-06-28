package main

import (
	"github.com/jteeuwen/glfw"
	"time"
)


/// returns milliseconds
func GetTime() int {
	return int(glfw.Time() * 1000)
}


type StopwatchStruct struct {
	Start time.Time
}

func Stopwatch() (sw StopwatchStruct) {
	sw.Start = time.Now()
	return
}

/// returns milliseconds
func (sw *StopwatchStruct) Stop() int64{
	return int64(time.Now().Sub(sw.Start)/1000000)
}

