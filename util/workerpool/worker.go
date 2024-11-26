package workerpool

type worker struct {
	taskChan chan *Task
	quitChan chan bool

	n int
}

func newWorker(n int, ch chan *Task) *worker {
	return &worker{
		n:        n,
		taskChan: ch,
		quitChan: make(chan bool),
	}
}

func (w *worker) startBackground() {
	for {
		select {
		case <-w.quitChan:
			return

		case task := <-w.taskChan:
			process(task)
		}
	}
}

func (w *worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}
