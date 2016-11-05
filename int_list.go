package fib

type intList struct {
	v    int
	tail *intList
}

func (self *intList) Pop() (int, *intList) {
	return self.v, self.tail
}

func (self *intList) Prepend(v int) *intList {
	return &intList{v, self}
}

func (self *intList) Value() int {
	return self.v
}
