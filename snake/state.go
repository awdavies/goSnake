package snake

type SnakeBody struct {
	next *SnakeBody
	x float64
	y float64
}

type SnakeState struct {
	dead bool
	head *SnakeBody
}

func NewSnakeState() *SnakeState {
	return &SnakeState {
		dead: false,
		head: &SnakeBody{nil, 0, 0},
	}	
}
