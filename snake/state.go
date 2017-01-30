package snake

type Direction int
const (
	kDirUp    Direction = 0
	kDirLeft  Direction = 1
	kDirRight Direction = 2
	kDirDown  Direction = 3
)

type Food struct {
	eaten bool

	// This is based on a grid.  Should be whole numbers.
	x float32
	y float32
}

type SnakeBody struct {
	Next *SnakeBody
	Tail *SnakeBody

	x float32
	y float32
}

type SnakeState struct {
	dead bool
	head *SnakeBody
	lastUpdate float64  // Last update timewise.
}

func NewSnakeState() (state *SnakeState) {
	snake_head := &SnakeBody{nil, nil, 2, 0}
	snake_head.Next = &SnakeBody{nil, nil, 1, 0}
	snake_head.Next.Next = &SnakeBody{nil, nil, 0, 0}
	snake_head.Tail = snake_head.Next.Next
	state = &SnakeState {
		dead: false,
		head: snake_head,
		lastUpdate: 0.0,
	}	
	return	
}

func (s *SnakeState) Update() {
	// TODO: Do some logic updates based on the last time this was called.
}
