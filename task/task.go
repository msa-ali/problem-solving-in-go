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

func (t *Task) Start() {
	t.setActiveDependency()
	// if Task has dependency
	if t.activeDependencyCount > 0 {
		for _, dependency := range t.dependencies {
			dependency.Subscribe(t.trackDependency)
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

func (t *Task) Subscribe(s subscriber) {
	t.subscribers = append(t.subscribers, s)
}

func (t *Task) publish() {
	for _, subscriber := range t.subscribers {
		subscriber()
	}
}

func (t *Task) done() {
	t.IsCompleted = true
	t.publish()
}
