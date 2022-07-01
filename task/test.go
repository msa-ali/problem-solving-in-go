package task

import (
	"fmt"
	"sync"
	"time"
)

func Test() {
	// Task 1 - no dependency
	task1 := New(nil, func(done func()) {
		time.AfterFunc(100*time.Millisecond, func() {
			fmt.Println("Task 1 Completed!")
			done()
		})
	})
	// Task 2 - no dependency
	task2 := New(nil, func(done func()) {
		time.AfterFunc(1500*time.Millisecond, func() {
			fmt.Println("Task 2 Completed!")
			done()
		})
	})
	// Task 3 - no dependency
	task3 := New(nil, func(done func()) {
		time.AfterFunc(1000*time.Millisecond, func() {
			fmt.Println("Task 3 Completed!")
			done()
		})
	})
	// Task 4 - have dependency
	task4 := New([]*Task{task1, task2}, func(done func()) {
		time.AfterFunc(100*time.Millisecond, func() {
			fmt.Println("Task 4 Completed!")
			done()
		})
	})

	task5 := New([]*Task{task3, task4}, func(done func()) {
		time.AfterFunc(1000*time.Millisecond, func() {
			fmt.Println("Task 5 Completed!")
			done()
		})
	})

	allDone := New([]*Task{task1, task2, task3, task4, task5}, func(done func()) {
		time.AfterFunc(1000*time.Millisecond, func() {
			fmt.Println("All Tasks are completed!")
			done()
		})
	})

	//  task3, task4, task5, allDone
	var wg sync.WaitGroup
	allTasks := []*Task{task1, task2, task3, task4, task5, allDone}
	for _, t := range allTasks {
		fmt.Print(fmt.Sprintf("Running Task %s", t.Id) + "\n")
		wg.Add(1)
		go func(t *Task) {
			t.Subscribe(func() {
				wg.Done()
			})
			t.Start()
		}(t)
	}

	wg.Wait()
}
