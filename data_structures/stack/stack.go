package stack

type Stack struct {
	Items []int
}

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() int {
	x, ss := s.Items[len(s.Items)-1], s.Items[:len(s.Items)-1]
	s.Items = ss
	return x
}

func (s *Stack) Peek() (int, bool) {
	if len(s.Items) > 0 {
		return s.Items[len(s.Items)-1], true
	}
	return 0, false
}

func NewStack() *Stack {
	return &Stack{
		Items: []int{},
	}
}
