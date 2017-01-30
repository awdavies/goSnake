package snake

type SnakeBody struct {
	next *SnakeBody
	x float64
	y float64
}

type SnakeState struct {
	dead bool
	head *SnakeBody
	lastUpdate float64  // Last update timewise.
}

func NewSnakeState() *SnakeState {
	return &SnakeState {
		dead: false,
		head: &SnakeBody{nil, 0, 0},
		lastUpdate: 0.0,
	}	
}

func (s *SnakeState) Update() {
	// TODO: Do some logic updates based on the last time this was called.
}
