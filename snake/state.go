package snake

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Direction int
const (
	kDirUp    Direction = 0
	kDirLeft  Direction = 1
	kDirRight Direction = 2
	kDirDown  Direction = 3
)

type GridState struct {
	Movement Direction
	ContainsSnake bool
}

type Food struct {
	eaten bool

	x int
	y int
}

type SnakeBody struct {
	Next *SnakeBody
	Tail *SnakeBody
	X int
	Y int
}

type SnakeState struct {
	dead bool
	head *SnakeBody
	lastUpdate float64  // Last update timewise.
	Grid [][]*GridState
	LastKey glfw.Key  // Last pressed important key.
}

func KeyToDirection(k glfw.Key) Direction {
	switch k {
		case glfw.KeyUp:
			return kDirUp
		case glfw.KeyDown:
			return kDirDown
		case glfw.KeyLeft:
			return kDirLeft
		case glfw.KeyRight:
			return kDirRight
	}
	return kDirRight
}

func DirectionToKey(d Direction) glfw.Key {
	switch d {
		case kDirUp:
			return glfw.KeyUp
		case kDirDown:
			return glfw.KeyDown
		case kDirLeft:
			return glfw.KeyLeft
		case kDirRight:
			return glfw.KeyRight
	}
	return glfw.KeyRight
}

func NewSnakeState() (state *SnakeState) {
	snake_head := &SnakeBody{nil, nil, 2, 0}
	snake_head.Next = &SnakeBody{nil, nil, 1, 0}
	snake_head.Next.Next = &SnakeBody{nil, nil, 0, 0}
	snake_head.Tail = snake_head.Next.Next
	state = &SnakeState {
		dead: false,
		head: snake_head,
	}
	// Allocate direction grid.
	state.Grid = make([][]*GridState, GridWidth)
	for i := 0; i < GridWidth; i++ {
		state.Grid[i] = make([]*GridState, GridHeight)
		for j := 0; j < GridHeight; j++ {
			state.Grid[i][j] = &GridState{kDirRight, false}
		}
	}
	node := snake_head
	for ; node != nil; node = node.Next {
		state.Grid[node.X][node.Y].ContainsSnake = true
	}
	state.lastUpdate = glfw.GetTime()
	return
}

func PollKeyPressHelper(w *glfw.Window, s *SnakeState, k glfw.Key) bool {
	// There are still some timing issues when switching directions rapidly.
	key_state := w.GetKey(k)
	if key_state == glfw.Press && s.LastKey != k {
		l := s.LastKey
		if (k == glfw.KeyUp && l == glfw.KeyDown) ||
			 (k == glfw.KeyDown && l == glfw.KeyUp) ||
			 (k == glfw.KeyLeft && l == glfw.KeyRight) ||
			 (k == glfw.KeyRight && l == glfw.KeyLeft) {
			return false
		}
		s.LastKey = k
		return true
	}
	return false
}

func (s *SnakeState) PollKeyPresses(w *glfw.Window) {
	if PollKeyPressHelper(w, s, glfw.KeyRight) {
		return
	}
	if PollKeyPressHelper(w, s, glfw.KeyLeft) {
		return
	}
	if PollKeyPressHelper(w, s, glfw.KeyUp) {
		return
	}
	if PollKeyPressHelper(w, s, glfw.KeyDown) {
		return
	}
}

func (s *SnakeState) Update(w *glfw.Window) {
	current_time := glfw.GetTime()
	elapsed_time := current_time - s.lastUpdate

	// TODO: Do something else instead?
	if s.dead {
		return
	}

	s.PollKeyPresses(w)
	if elapsed_time > 0.1 {
		node := s.head
		// Figure out the next move for the head node.
		s.Grid[node.X][node.Y].Movement = KeyToDirection(s.LastKey)
		for ; node != nil; node = node.Next {
			s.Grid[node.X][node.Y].ContainsSnake = false
			var new_x, new_y int
			switch move := s.Grid[node.X][node.Y].Movement; move {
				case kDirRight: {
					new_x = node.X + 1
					new_y = node.Y
				}
				case kDirLeft: {
					new_x = node.X - 1
					new_y = node.Y
				}
				case kDirUp: {
					new_x = node.X
					new_y = node.Y - 1
				}
				case kDirDown: {
					new_x = node.X
					new_y = node.Y + 1
				}
			}
			// We're dead if we're outside the grid.
			if new_x >= GridWidth || new_y >= GridHeight ||
				 new_y < 0 || new_x < 0 {
				s.dead = true
				return
			}
			// If we've hit another snake part... we're dead.
			if s.Grid[new_x][new_y].ContainsSnake {
				s.dead = true
				return
			}
			s.Grid[new_x][new_y].ContainsSnake = true
			node.X = new_x
			node.Y = new_y
		}

		s.lastUpdate = current_time
	}
}
