package queue

import (
	"fmt"
	"sync"
)

// TaskQueue manages a sequential task execution queue.
type TaskQueue struct {
	tasks    chan func()
	stopChan chan struct{}
	wg       sync.WaitGroup
}

// NewTaskQueue creates a new TaskQueue with a buffered channel for tasks.
func NewTaskQueue(bufferSize int) *TaskQueue {
	q := &TaskQueue{
		tasks:    make(chan func(), bufferSize),
		stopChan: make(chan struct{}),
	}
	go q.start() // Start processing tasks.
	return q
}

// start processes tasks sequentially until the queue is stopped.
func (q *TaskQueue) start() {
	for {
		select {
		case task, ok := <-q.tasks:
			if !ok {
				return // Channel closed, exit the loop.
			}
			q.wg.Add(1) // Increment the WaitGroup counter.
			func() {
				defer q.wg.Done() // Ensure Done is called after task completion.
				task()
			}()
		case <-q.stopChan:
			return // Stop signal received, exit the loop.
		}
	}
}

// AddTask adds a task to the queue for sequential execution.
func (q *TaskQueue) AddTask(task func()) error {
	select {
	case q.tasks <- task:
		return nil // Task successfully added.
	default:
		return fmt.Errorf("task queue is full")
	}
}

// Stop gracefully shuts down the queue by waiting for tasks to complete.
func (q *TaskQueue) Stop() {
	close(q.stopChan) // Signal the queue to stop.

	// Wait for all ongoing tasks to finish.
	q.wg.Wait()

	// Optional: Drain and close the task channel.
	for len(q.tasks) > 0 {
		<-q.tasks
	}
	close(q.tasks)
}
