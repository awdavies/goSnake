package main

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	"./snake"
)

const kWindowTitle = "Snake!"

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(snake.Width, snake.Height, 
																	 kWindowTitle, nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.MakeContextCurrent()

	// Initialize Glow.
	if err := gl.Init(); err != nil {
		panic(err)
	}
	window.SetKeyCallback(snake.SnakeKeyCallback)
	glfw.SwapInterval(1)  // Turn on VSync

	// Main loop.
	state := snake.NewSnakeState()
	for !window.ShouldClose() {
		state.Update(window)
		state.Draw()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
