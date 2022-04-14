package Chapter01Foundation

/*func main() {
	stackOfStrings := NewFixedCapStackOfStrings(3)
	for i:=0; i < 4; i++ {
		if ok := stackOfStrings.Push("abc" + strconv.Itoa(i)); !ok {
			fmt.Println("error")
			continue
		}
		fmt.Println("success")
	}
	for i := 0; i < 4; i++ {
		fmt.Println(stackOfStrings.Pop())
	}
}*/

type FixedCapStackOfStrings struct {
	Strs []string
	N int
}

func NewFixedCapStackOfStrings(capacity int) *FixedCapStackOfStrings {
	return &FixedCapStackOfStrings{
		N: capacity,
		Strs: make([]string, 0),
	}
}

func (s *FixedCapStackOfStrings) IsEmpty() bool {
	return len(s.Strs) == 0
}

func (s *FixedCapStackOfStrings) Size() int {
	return len(s.Strs)
}

func (s *FixedCapStackOfStrings) Push(item string) bool {
	if s.Size() >= s.N {
		return false
	}
	s.Strs = append(s.Strs, item)
	return true
}

func (s *FixedCapStackOfStrings) Pop() string {
	if s.Size() == 0 {
		return ""
	}
	res := s.Strs[s.Size()-1]
	s.Strs = s.Strs[:s.Size()-1]
	return res
}