package fib

import "testing"

func TestintList(t *testing.T) {
	firstNode := intList{3, nil}
	secondNode := firstNode.Prepend(2)
	thirdNode := secondNode.Prepend(3)
	v1, t1 := thirdNode.Pop()
	v2, t2 := t1.Pop()
	v3, t3 := t2.Pop()
	if v1 != 3 && v2 != 2 && v3 != 1 && t3 != nil {
		t.Error("List is broken")
	}
}
