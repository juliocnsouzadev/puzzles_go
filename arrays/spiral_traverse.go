package arrays

const (
	leftToRight Direction = "leftToRight"
	topToBottom Direction = "topToBottom"
	rightToLeft Direction = "rightToLeft"
	bottomToTop Direction = "bottomToTop"
)

func SpiralTraverse(array [][]int) []int {
	height := len(array)
	width := len(array[0])
	result := []int{}

	state := NewState(0, width-1, 0, height-1, leftToRight)

	for state.CanChangeState() {
		switch state.direction {
		case leftToRight:
			i := state.left
			for i <= state.right {
				result = append(result, array[state.top][i])
				i++
			}
			state.NextState(topToBottom)
			state.IncreaseTop()

		case topToBottom:
			i := state.top
			for i <= state.bottom {
				result = append(result, array[i][state.right])
				i++
			}
			state.NextState(rightToLeft)
			state.DecreaseRight()

		case rightToLeft:
			i := state.right
			for i >= state.left {
				result = append(result, array[state.bottom][i])
				i--
			}
			state.NextState(bottomToTop)
			state.DecreaseBottom()

		case bottomToTop:
			i := state.bottom
			for i >= state.top {
				result = append(result, array[i][state.left])
				i--
			}
			state.NextState(leftToRight)
			state.IncreaseLeft()

		}
	}

	return result
}

type Direction string

type State struct {
	left      int
	right     int
	top       int
	bottom    int
	direction Direction
}

func NewState(left, right, top, bottom int, direction Direction) *State {
	return &State{
		left:      left,
		right:     right,
		top:       top,
		bottom:    bottom,
		direction: direction,
	}
}

func (s *State) NextState(direction Direction) {
	s.direction = direction
}

func (s *State) CanChangeState() bool {
	return s.left <= s.right && s.top <= s.bottom
}

func (s *State) IncreaseTop() {
	s.top++
}

func (s *State) DecreaseRight() {
	s.right--
}

func (s *State) DecreaseBottom() {
	s.bottom--
}

func (s *State) IncreaseLeft() {
	s.left++
}
