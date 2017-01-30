package main

import (
	"runtime"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const kWindowTitle = "Snake!"

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, kWindowTitle, nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Main event loop.
	window.MakeContextCurrent()
	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
