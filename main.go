package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Altamashattari/problem-solving/task"
)

func main() {
	// Task 1
	task1 := task.Task{
		Task: func(done func()) {
			time.AfterFunc(100*time.Millisecond, func() {
				fmt.Println("Task 1 Completed!")
				done()
			})
		},
	}

	task2 := task.Task{
		Task: func(done func()) {
			time.AfterFunc(1500*time.Millisecond, func() {
				fmt.Println("Task 2 Completed!")
				done()
			})
		},
	}

	task3 := task.Task{
		Task: func(done func()) {
			time.AfterFunc(1000*time.Millisecond, func() {
				fmt.Println("Task 3 Completed!")
				done()
			})
		},
	}

	task4 := task.Task{
		Dependencies: []task.Task{task1, task2},
		Task: func(done func()) {
			time.AfterFunc(1000*time.Millisecond, func() {
				fmt.Println("Task 4 Completed!")
				done()
			})
		},
	}

	task5 := task.Task{
		Dependencies: []task.Task{task3, task4},
		Task: func(done func()) {
			time.AfterFunc(1000*time.Millisecond, func() {
				fmt.Println("Task 5 Completed!")
				done()
			})
		},
	}

	allDone := task.Task{
		Dependencies: []task.Task{task1, task2, task3, task4, task5},
		Task: func(done func()) {
			time.AfterFunc(1000*time.Millisecond, func() {
				fmt.Println("All Tasks are completed!")
				done()
			})
		},
	}

	var wg sync.WaitGroup
	allTasks := []task.Task{task1, task2, task3, task4, task5, allDone}

	for _, t := range allTasks {
		wg.Add(1)
		go func(t *task.Task) {
			t.Subscribe(func() {
				wg.Done()
			})
			t.Start()
		}(&t)
	}

	wg.Wait()

}
