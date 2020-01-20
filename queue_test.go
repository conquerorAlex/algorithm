package algorithm

import (
	"testing"
)

func TestEnqueue(t *testing.T){
	t.Log("testQueue")
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Print()
}

func TestDequeue(t *testing.T){
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Print()
	q.Dequeue()
	q.Print()
}