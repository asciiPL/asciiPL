package screen

type Stack []*Page

func NewStack() *Stack {
	s := make(Stack, 0)
	return &s
}

func (s *Stack) Push(value *Page) {
	*s = append(*s, value)
}

func (s *Stack) Pop() *Page {
	if s.IsEmpty() {
		return nil
	}
	lastIndex := len(*s) - 1
	value := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return value
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Peek() *Page {
	if s.IsEmpty() {
		return nil
	}
	lastIndex := len(*s) - 1
	return (*s)[lastIndex]
}
