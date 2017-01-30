package snake

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func SnakeKeyCallback(w *glfw.Window, 
			  						 key glfw.Key, 
										 scancode int, 
										 action glfw.Action, 
										 mods glfw.ModifierKey) {
	if (key == glfw.KeyEscape && action == glfw.Press) {
		w.SetShouldClose(true)
	}
}
