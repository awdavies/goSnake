package snake

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func (s *SnakeBody) Draw(w *glfw.Window) {
	// TODO: Draw some stuff.
}

func (s *SnakeState) Draw(w *glfw.Window) {
	// TODO: Need some logic on what to draw, like if we lose, etc.
	s.head.Draw(w)
}
