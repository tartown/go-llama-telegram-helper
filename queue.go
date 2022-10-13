
package main

import "sync"

type TaskQueue struct {
	mu sync.Mutex
	tasks []*Task
	users map[int64]*Task
	Limit int
	Count int
}

func NewTaskQueue(limit int) *TaskQueue {
	return &TaskQueue{
		tasks: make([]*Task, 0),
		users: make(map[int64]*Task, 0),
		Limit: limit,