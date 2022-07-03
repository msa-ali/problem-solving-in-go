package worklist

type Entry struct {
	Path string
}

type Worklist struct {
	jobs chan Entry
}

func (w *Worklist) Add(work Entry) {
	w.jobs <- work
}

func (w *Worklist) Next() Entry {
	j := <-w.jobs
	return j
}

func New(bufSize int) Worklist {
	return Worklist{make(chan Entry, bufSize)}
}

func NewJob(path string) Entry {
	return Entry{path}
}
func (w *Worklist) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Add(Entry{""})
	}
}
