package task

type taskFunc func(done func())

type subscriber func()

type Task struct {
	IsCompleted  bool
	Task         taskFunc
	Dependencies []Task

	activeDependencyCount int
	subscribers           []subscriber
}

func (t *Task) Start() {
	t.setActiveDependency()

	// if Task has dependency
	if t.activeDependencyCount > 0 {
		for _, dependency := range t.Dependencies {
			dependency.Subscribe(dependency.trackDependency)
		}
		return
	}

	// if no dependency, execute actual task
	t.executeTask()
}

func (t *Task) executeTask() {
	t.Task(t.done)
}

func (t *Task) setActiveDependency() {
	var updatedDependencyList []Task

	for _, dependency := range t.Dependencies {
		if dependency.IsCompleted {
			updatedDependencyList = append(updatedDependencyList, dependency)
		}
	}
	t.activeDependencyCount = len(updatedDependencyList)
	t.Dependencies = updatedDependencyList
}

func (t *Task) trackDependency() {
	t.activeDependencyCount--
	if t.activeDependencyCount == 0 {
		t.executeTask()
	}
}

func (t *Task) Subscribe(s subscriber) {
	t.subscribers = append(t.subscribers, s)
}

func (t *Task) done() {
	t.IsCompleted = true
	for _, subscriber := range t.subscribers {
		subscriber()
	}
}
