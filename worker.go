package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A Worker is the interface Wait and Do methods. Each worker lives in its own
// goroutine.
//
// Wait method waits task, from jobs channel, to be pulled off of queue. Do method
// runs the task.
type Worker interface {
	Wait(jobs <-chan *Task, results chan<- *TaskResult)
	Do(task *Task) *TaskResult
	GetWorkerID() string
}

var workerFactory = map[TaskType]func(id string) Worker{
	TaskSendEmail:        NewMailSender,
	TaskGenerateThumbail: NewThumbnailGenerator,
}

// MailSender specifies worker that will be sending email.
type MailSender struct {
	WorkerID string
	Worker
}

// NewMailSender creates MailSender worker.
func NewMailSender(id string) Worker {
	fmt.Printf("Starting worker '%s'\n", id)
	return &MailSender{
		WorkerID: id,
	}
}

// Wait waits for incoming email to be sent.
func (ms *MailSender) Wait(jobs <-chan *Task, results chan<- *TaskResult) {
	var task *Task
	for {
		// Get a task from the jobs queue.
		task = <-jobs
		if task == nil {
			continue
		}
		if task.Type != TaskSendEmail {
			continue
		}

		// Do the task.
		results <- ms.Do(task)
	}
}

// Do sends email.
//
// Since this is a dummy MailSender, it directly pipes the result to results
// channel.
func (ms *MailSender) Do(task *Task) *TaskResult {
	result := &TaskResult{
		Task:    task,
		Worker:  ms,
		Started: time.Now(),
	}
	fmt.Printf("%s: taking %s\n", ms.WorkerID, task.TaskID)

	result.Result = doProcessing()
	result.Completed = time.Now()

	return result
}

func (ms *MailSender) GetWorkerID() string {
	return ms.WorkerID
}

// A ThumbnailGenerator specifies worker that will be generating thumbnail.
type ThumbnailGenerator struct {
	WorkerID string
	Worker
}

// NewThumbnailGenerator creates ThumbnailGenerator worker.
func NewThumbnailGenerator(id string) Worker {
	fmt.Printf("Starting worker '%s'\n", id)
	return &ThumbnailGenerator{
		WorkerID: id,
	}
}

// Wait waits for incoming image that needs thumbnail to be generated.
func (tg *ThumbnailGenerator) Wait(jobs <-chan *Task, results chan<- *TaskResult) {
	var task *Task
	for {
		// Get a task from the jobs queue.
		task = <-jobs
		if task == nil {
			continue
		}
		if task.Type != TaskGenerateThumbail {
			continue
		}

		// Do the task.
		results <- tg.Do(task)
	}
}

// Do generates thumbnail.
func (tg *ThumbnailGenerator) Do(task *Task) *TaskResult {
	result := &TaskResult{
		Task:    task,
		Worker:  tg,
		Started: time.Now(),
	}
	fmt.Printf("%s: taking %s\n", tg.WorkerID, task.TaskID)

	result.Result = doProcessing()
	result.Completed = time.Now()

	return result
}

func (tg *ThumbnailGenerator) GetWorkerID() string {
	return tg.WorkerID
}

// Do dummy processing until random time and returns random result.
func doProcessing() bool {
	rand.Seed(time.Now().UTC().UnixNano())

	d := time.Duration(rand.Intn(3)) * time.Second
	<-time.After(d)

	return rand.Intn(2) == 1
}
