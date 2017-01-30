package snake

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func (s *SnakeBody) Draw(w *glfw.Window) {
	// TODO: Draw some stuff.
}

func (s *SnakeState) Draw(w *glfw.Window) {
	// TODO: Need some logic on what to draw, like if we lose, etc.
	s.head.Draw(w)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Begin(gl.QUADS)
	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex2f(-0.5, -0.5);
	gl.Vertex2f( 0.5, -0.5);
	gl.Vertex2f( 0.5,  0.5);
	gl.Vertex2f(-0.5,  0.5);
	gl.End();
	gl.Flush();
}
