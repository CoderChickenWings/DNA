package vm


type OpStack struct {
	Element []*StackItem
}

func NewOpStack() *OpStack {
	var stack OpStack
	e := make([]*StackItem,0)
	stack.Element = e
	return &stack;
}

func  (s *OpStack) Push (data *StackItem ) {
	// Stack add
	//stack.list.PushBack(data)
	s.Element = append(s.Element,data)
}

func  (s *OpStack) Peek() *StackItem{
	//return stack.list.Back()
	len := len(s.Element)
	if(len == 0) {return nil}

	return s.Element[len-1]
}

func  (s *OpStack) RemoveFirst() *StackItem {
	len := len(s.Element)
	if(len == 0) {return nil}

	e := s.Element[len-1]

	if  len-1 == 0 {
		s.Element = nil
	} else {
		s.Element = s.Element[:len-1]
	}

	return e
}

func  (s *OpStack) Pop() *StackItem{

	return s.RemoveFirst()
}


func  (s *OpStack) Count() int{
	return len(s.Element)
}



