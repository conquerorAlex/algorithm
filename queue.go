package algorithm

import "fmt"

type Item interface {}

type Queuer interface{
	Enqueue(t Item)
	Dequeue() *Item
	Size() int
	IsEmpty() bool
	Print()
}

type queue struct{
	items []Item
}

func NewQueue() Queuer{
	return &queue{}
}

func (q *queue)Print(){
	fmt.Println(q.items)
}

func (q *queue)Enqueue(t Item){
	if q == nil{
		fmt.Println("queue nil")
		return
	}

	q.items = append(q.items, t)
}

func (q *queue)Dequeue() *Item{
	if q == nil{
		fmt.Println("queue nil")
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return &item
}

func (q *queue)Size()int{
	if q == nil {
		fmt.Println("queue nil")
		return 0
	}

	return len(q.items)
}

func (q *queue)IsEmpty()bool{
	if q == nil {
		fmt.Println("queue empty")
		return true
	}

	return len(q.items) == 0
}