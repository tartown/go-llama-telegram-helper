package main

import "fmt"

type Task struct {
	UserID         int64
	MessageID      int
	AnnounceID     int
	Question       string
	Stopped        bool
	Stop        chan bool
}

func (t *Task) 