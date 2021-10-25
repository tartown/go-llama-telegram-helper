package main

import "errors"

var ErrQueueEmpty = errors.New("queue is empty")
var ErrOnePerUser = errors.New("user already app