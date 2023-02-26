package main

type TaskStatus int

const (
	_ TaskStatus = iota
	Pending
	Completed
)

type Task struct {
	TaskId int32
	Title  string
	Status TaskStatus
}
