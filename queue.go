package main

import (
	"fmt"
)

// A Queue specifies jobs queue which wraps taskType, the channel jobs that
// receives task with type taskType, and the results channel.
type Queue struct {
	taskType TaskType
	jobs     chan<- *Task
	results  <-chan *TaskResult
}

// Starts n workers that consume Task type t.
func NewQueue(n int, t TaskType) *Queue {
	fmt.Printf("Running '%s' queue with %d worker(s)\n", taskNames[t], n)

	jobs := make(chan *Task)
	results := make(chan *TaskResult)
	queue := &Queue{
		taskType: t,
		jobs:     jobs,
		results:  results,
	}

	for i := 0; i < n; i++ {
		worker := workerFactory[t](fmt.Sprintf("Worker#%s#%d", taskNames[t], i+1))

		go worker.Wait(jobs, results)
	}

	return queue
}
