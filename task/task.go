package task

import (
	"fmt"
	"strconv"
)

type taskFunc func(done func())

type subscriber func()

type Task struct {
	Id                    string
	IsCompleted           bool
	task                  taskFunc
	dependencies          []*Task
	activeDependencyCount int
	subscribers           []subscriber
	// send-only channel list
	channels []chan<- struct{}
}

var id int = 0

func New(dependencies []*Task, t taskFunc) *Task {
	id++
	return &Task{
		Id:           strconv.Itoa(id),
		task:         t,
		dependencies: dependencies,
	}
}

func (t *Task) Start(useChan bool) {
	t.setActiveDependency()
	// if Task has dependency
	if t.activeDependencyCount > 0 {
		for _, dependency := range t.dependencies {
			if useChan {
				ch := dependency.Subscribe(nil, true)
				go func(ch <-chan struct{}) {
					<-ch
					t.trackDependency()
				}(ch)
			} else {
				dependency.Subscribe(t.trackDependency, false)
			}
		}
		fmt.Printf("Task %s has %d dependencies. Waiting for them to get resolved...\n", t.Id, t.activeDependencyCount)
		return
	}

	// if no dependency, execute actual task
	t.executeTask()
}

func (t *Task) executeTask() {
	t.task(t.done)
}

func (t *Task) setActiveDependency() {
	var updatedDependencyList []*Task

	for _, dependency := range t.dependencies {
		if !dependency.IsCompleted {
			updatedDependencyList = append(updatedDependencyList, dependency)
		}
	}
	t.activeDependencyCount = len(updatedDependencyList)
	t.dependencies = updatedDependencyList
}

func (t *Task) trackDependency() {
	t.activeDependencyCount -= 1
	if t.activeDependencyCount == 0 {
		t.executeTask()
	}
}

func (t *Task) Subscribe(s subscriber, useChan bool) <-chan struct{} {
	if useChan {
		ch := make(chan struct{})
		t.channels = append(t.channels, ch)
		return ch
	}
	t.subscribers = append(t.subscribers, s)
	return nil
}

func (t *Task) publish() {
	for _, subscriber := range t.subscribers {
		subscriber()
	}
	if len(t.channels) > 0 {
		for _, ch := range t.channels {
			ch <- struct{}{}
			close(ch)
		}
	}
}

func (t *Task) done() {
	t.IsCompleted = true
	t.publish()
}
