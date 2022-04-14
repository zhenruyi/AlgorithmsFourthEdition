package Chapter01Foundation

/*func main() {
	stack := NewFixedCapStack(3)
	stack.Push(123)
	stack.Push("aaa")
	stack.Push(123.123)
	stack.Push("bbb")
	for i:=0; i < 3; i++ {
		val := stack.Pop()
		fmt.Printf("%T:%v\n", val, val)
	}
}*/

type FixedCapStack struct {
	data []interface{}
	capa int
}

func NewFixedCapStack(capacity int) *FixedCapStack {
	return &FixedCapStack{
		capa: capacity,
		data: make([]interface{}, 0),
	}
}

func (s *FixedCapStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *FixedCapStack) Size() int {
	return len(s.data)
}

func (s *FixedCapStack) Push(d interface{}) bool {
	if len(s.data) >= s.capa {
		return false
	}
	s.data = append(s.data, d)
	return true
}

func (s *FixedCapStack) Pop() interface{} {
	if s.Size() == 0 {
		return nil
	}
	res := s.data[s.Size()-1]
	s.data = s.data[:s.Size()-1]
	return res
}

func (s *FixedCapStack) Resize(size int) {
	s.capa *= 2
}
