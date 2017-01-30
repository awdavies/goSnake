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

func DrawSquareAt(x int, y int, r, g, b float32) {
	gl.Begin(gl.QUADS)
	gl.Color3f(r, g, b)  // Snakes are green, obviously!
	gl.Vertex2f(-1.0 + XDrawDelta * float32(x),
							 1.0 - YDrawDelta * float32(y + 1))
	gl.Vertex2f(-1.0 + XDrawDelta * float32(x + 1),
							 1.0 - YDrawDelta * float32(y + 1))
	gl.Vertex2f(-1.0 + XDrawDelta * float32(x + 1),
						   1.0 - YDrawDelta * float32(y))
	gl.Vertex2f(-1.0 + XDrawDelta * float32(x),
							 1.0 - YDrawDelta * float32(y))
	gl.End();
}

func (s *SnakeBody) Draw() {
	if s == nil {
		return
	}
	node := s
	for node != nil {
		DrawSquareAt(node.X, node.Y, 0, 1.0, 0)
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
