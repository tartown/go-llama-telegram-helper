
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
	}
}


// Get task by UserID and its count in queue
func (q *TaskQueue) Load(userId int64) (*Task, int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for n, task := range q.tasks {
		if task.UserID == userId {
			return task, n
		}
	}

	return nil, -1
}


func (q *TaskQueue) Enqueue(task *Task) (int, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	t, exists := q.users[task.UserID]
	if exists {