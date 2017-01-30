package snake

import (
	"github.com/go-gl/gl/v2.1/gl"
)

const (
	Width int = 640
	Height int = 480

	GridWidth = 64
	GridHeight = 48

	// The delta for creating a grid of 10px by 10px squares.
	XDrawDelta float32 = 0.03125
	YDrawDelta float32 = 0.041667
)

func (s *SnakeBody) Draw() {
	if s == nil {
		return
	}
	node := s
	for node != nil {
		gl.Begin(gl.QUADS)
		gl.Color3f(0.0, 1.0, 0.0)  // Snakes are green, obviously!
		gl.Vertex2f(-1.0 + XDrawDelta * float32(node.X),  
								 1.0 - YDrawDelta * float32(node.Y + 1))
		gl.Vertex2f(-1.0 + XDrawDelta * float32(node.X + 1), 
								 1.0 - YDrawDelta * float32(node.Y + 1))
		gl.Vertex2f(-1.0 + XDrawDelta * float32(node.X + 1), 
							   1.0 - YDrawDelta * float32(node.Y))
		gl.Vertex2f(-1.0 + XDrawDelta * float32(node.X), 
								 1.0 - YDrawDelta * float32(node.Y))
		gl.End();
		node = node.Next
	}
}

func (s *SnakeState) Draw() {
	// TODO: Need some logic on what to draw, like if we lose, etc.
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	s.head.Draw()
	gl.Flush();
}
