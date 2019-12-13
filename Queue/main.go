package main

import "fmt"

type queue struct {
	queue []string
}

func (q *queue) enqueue(v string) {
	q.queue = append(q.queue, v)
	fmt.Println(q.queue)
}

func (q *queue) dequeue() {
	q.queue = q.queue[1:]
	fmt.Println(q.queue)
}

func main() {
	q := queue{queue:make([]string, 0)}
	q.enqueue("1")
	q.enqueue("2")
	q.dequeue()
}
