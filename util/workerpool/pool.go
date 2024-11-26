package workerpool

type Pool struct {
	collector   chan *Task
	workers     []*worker
	workerCount int
}

func NewPool(workerCount, collectorCount int) *Pool {
	return &Pool{
		workerCount: workerCount,
		workers:     make([]*worker, workerCount),
		collector:   make(chan *Task, collectorCount),
	}
}

func (p *Pool) AddTask(t *Task) {
	p.collector <- t
}

func (p *Pool) RunBackground() {
	for i := 0; i < p.workerCount; i++ {
		w := newWorker(i, p.collector)
		p.workers[i] = w

		go w.startBackground()
	}
}

func (p *Pool) Stop() {
	for i := 0; i < p.workerCount; i++ {
		p.workers[i].stop()
	}
}
