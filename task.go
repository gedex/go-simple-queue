package main

import (
	"time"
)

// A TaskType specifies type of task.
type TaskType int

const (
	TaskSendEmail TaskType = iota
	TaskGenerateThumbail
)

var taskNames = []string{
	"SendEmail",
	"GenerateThumbnail",
}

type TaskData map[string]string

type Task struct {
	TaskID string
	Type   TaskType
	Data   TaskData
}

type TaskResult struct {
	Task      *Task
	Worker    Worker
	Started   time.Time
	Completed time.Time
	Result    bool
}
