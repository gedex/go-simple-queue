package main

import (
	"fmt"
	"sync"
	"time"
)

type itemID struct {
	id int
	sync.Mutex
}

var (
	mailJobId = new(itemID)
	imgJobId  = new(itemID)
)

func main() {
	// Creates mail queue with 4 workers. Worker job is to send email that's
	// pull-off of queue.
	mailQueue := NewQueue(4, TaskSendEmail)
	go publishJobsEvery(mailQueue, time.Second)

	// Creates image queue with 3 workers. Worker job is to generate thumbnail
	// from image information that's pull-off of queue.
	imgQueue := NewQueue(3, TaskGenerateThumbail)
	go publishJobsEvery(imgQueue, time.Second*2)

	for {
		select {
		case result1 := <-mailQueue.results:
			showResult(result1)
		case result2 := <-imgQueue.results:
			showResult(result2)
		}
	}
}

// Publish jobs periodically to queue.
func publishJobsEvery(q *Queue, d time.Duration) {
	time.AfterFunc(d, func() {
		publishJobs(q)
		publishJobsEvery(q, d)
	})
}

// Publish jobs to queue.
func publishJobs(q *Queue) {
	switch q.taskType {
	case TaskSendEmail:
		q.jobs <- &Task{
			TaskID: fmt.Sprintf("Job#%s#%d", taskNames[q.taskType], mailJobId.getLastID()),
			Type:   TaskSendEmail,
			Data: TaskData{
				"from":  "admin@example.com",
				"to":    "user@somewhere.com",
				"title": "Test Email",
			},
		}
	case TaskGenerateThumbail:
		q.jobs <- &Task{
			TaskID: fmt.Sprintf("Job#%s#%d", taskNames[q.taskType], imgJobId.getLastID()),
			Type:   TaskGenerateThumbail,
			Data: TaskData{
				"img":              "/path/to/img",
				"width":            "1024",
				"height":           "768",
				"thumbnail_width":  "48",
				"thumbnail_height": "48",
			},
		}
	}
}

func showResult(tr *TaskResult) {
	var result string
	if tr.Result {
		result = "SUCCESS"
	} else {
		result = "FAIL"
	}
	fmt.Printf("%s done by %s [%s]\n", tr.Task.TaskID, tr.Worker.GetWorkerID(), result)
}

func (i *itemID) getLastID() int {
	i.Lock()
	defer i.Unlock()

	i.id += 1
	return i.id
}
