package utils

type Queue []interface{}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(element interface{}) {
	*q = append(*q, element)
}

func (q *Queue) Dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
