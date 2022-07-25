package circelqueue

import "errors"

const QueueSize int = 5

type CircleQueue struct {
	data  [QueueSize]interface{}
	front int
	rear  int
}

func InitQueue() *CircleQueue {
	q := &CircleQueue{}
	q.front = 0
	q.rear = 0
	return q
}

func DeQueue(q *CircleQueue) (interface{}, error) {
	if q.rear == q.front {
		return nil, errors.New("空队列")
	}
	res := q.data[q.front]
	q.data[q.front] = 0
	q.front = (q.front + 1) % QueueSize
	return res, nil
}
func EnQueue(q *CircleQueue, data interface{}) error {
	if (q.rear+1)%QueueSize == q.front%QueueSize {
		return errors.New("队列已满")
	}
	q.data[q.rear] = data
	q.rear = (q.rear + 1) % QueueSize

	return nil
}
func QueueLength(q *CircleQueue) int {
	return (q.rear - q.front + QueueSize) % QueueSize
}
