package snake

import (
	"math/rand"
	"github.com/go-gl/glfw/v3.2/glfw"
	"time"
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
	IsEaten bool

	X int
	Y int
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
	InputQueue []glfw.Key
	KeyPressed []bool
	NextFood *Food
	GrowLength int
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

func (s *SnakeState) GenerateFoodCoords() (int, int) {
	food_x := rand.Int() % GridWidth
	food_y := rand.Int() % GridHeight
	for s.Grid[food_x][food_y].ContainsSnake {
		food_x = rand.Int() % GridWidth
		food_y = rand.Int() % GridHeight
	}
	return food_x, food_y
}

func NewSnakeState() (state *SnakeState) {
	rand.Seed(time.Now().UTC().UnixNano())  // Generates new seed for food spawn.

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
	state.InputQueue = make([]glfw.Key, 0, 1024)
	// InputQueue always has one element.
	state.InputQueue = append(state.InputQueue, glfw.KeyRight)
	state.KeyPressed = make([]bool, 4)
	food_x, food_y := state.GenerateFoodCoords()
	state.NextFood = &Food{IsEaten: false, X: food_x, Y: food_y}
	state.GrowLength = 0
	return
}

func PollKeyPressHelper(w *glfw.Window, s *SnakeState, k glfw.Key) bool {
	// There are still some timing issues when switching directions rapidly.
	key_state := w.GetKey(k)
	index := KeyToDirection(k)
	last_key := s.InputQueue[len(s.InputQueue) - 1]
	if key_state == glfw.Press && s.KeyPressed[index] {
		return false
	}
	if key_state == glfw.Release && s.KeyPressed[index] {
		s.KeyPressed[index] = false
		return false
	}
	if key_state == glfw.Press && last_key != k {
		if (k == glfw.KeyUp && last_key == glfw.KeyDown) ||
			 (k == glfw.KeyDown && last_key == glfw.KeyUp) ||
			 (k == glfw.KeyLeft && last_key == glfw.KeyRight) ||
			 (k == glfw.KeyRight && last_key == glfw.KeyLeft) {
			return false
		}
		s.KeyPressed[index] = true
		s.InputQueue = append(s.InputQueue, k)
		return true
	}
	return false
}

func (s *SnakeState) PollKeyPresses(w *glfw.Window) {
	PollKeyPressHelper(w, s, glfw.KeyRight)
	PollKeyPressHelper(w, s, glfw.KeyLeft)
	PollKeyPressHelper(w, s, glfw.KeyUp)
	PollKeyPressHelper(w, s, glfw.KeyDown)
}

func (s *SnakeState) Update(w *glfw.Window) {
	current_time := glfw.GetTime()
	elapsed_time := current_time - s.lastUpdate
	// TODO: Do something else instead?  A "you died" screen?
	if s.dead {
		return
	}
	s.PollKeyPresses(w)
	if elapsed_time > 0.1 {
		if s.NextFood.IsEaten {
			s.NextFood.X, s.NextFood.Y = s.GenerateFoodCoords()
			s.NextFood.IsEaten = false
		}
		node := s.head
		// Figure out the next move for the head node.
		next_move := s.InputQueue[0]
		if len(s.InputQueue) > 1 {
			s.InputQueue = s.InputQueue[1:]  // Pops head from queue.
		}
		s.Grid[s.head.X][s.head.Y].Movement = KeyToDirection(next_move)
		// Growth check occurs before node iteration, as the snake techically isn't
		// moving when growing.
		growing := s.GrowLength > 0
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
			if new_x == s.NextFood.X && new_y == s.NextFood.Y {
				s.NextFood.IsEaten = true
				s.GrowLength += 4
			}
			s.Grid[new_x][new_y].ContainsSnake = true
			if growing {
				s.GrowLength -= 1
				new_snake_part := &SnakeBody{node.Next, node.Tail, node.X, node.Y}
				node.Next = new_snake_part
				s.Grid[node.X][node.Y].ContainsSnake = true
				node.X = new_x
				node.Y = new_y
				break
			}
			node.X = new_x
			node.Y = new_y
		}
		s.lastUpdate = current_time
	}
}
