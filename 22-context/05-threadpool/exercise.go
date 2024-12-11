package threadpool

import (
	"context"
	"log"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

type Runnable interface {
	Run(context.Context) error
}

type ThreadPool interface {
	Run(Runnable)
	Close()
}

func NewThreadPool(n int) (ThreadPool, chan error) {
	tasks := make(chan Runnable, n) // Buffered to prevent blocking
	errChan := make(chan error, 100)

	ctx, cancel := context.WithCancel(context.Background())

	pool := &threadPool{
		tasks:      tasks,
		errChan:    errChan,
		cancelFunc: cancel,
		closed:     false,
	}

	for i := 0; i < n; i++ {
		pool.waitGroup.Add(1)
		go pool.worker(ctx)
	}

	return pool, errChan
}

// threadPool implements the ThreadPool interface
type threadPool struct {
	tasks      chan Runnable
	errChan    chan error
	cancelFunc context.CancelFunc
	waitGroup  sync.WaitGroup
	once       sync.Once
	lock       sync.Mutex
	closed     bool
}

// worker processes tasks from the task channel
func (tp *threadPool) worker(ctx context.Context) {
	defer tp.waitGroup.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case task, ok := <-tp.tasks:
			if !ok {
				return
			}
			err := task.Run(ctx)
			if err != nil {
				select {
				case tp.errChan <- err:
				default:
					log.Printf("Error channel full, dropping error: %v\n", err)
				}
			}
		}
	}
}

// Run submits a Runnable task to the thread pool
func (tp *threadPool) Run(task Runnable) {
	tp.lock.Lock()
	defer tp.lock.Unlock()

	if tp.closed {
		log.Println("Failed to submit task: thread pool is closed")
		return
	}

	select {
	case tp.tasks <- task:
	default:
		log.Println("Failed to submit task: task channel is full")
	}
}

// Close gracefully shuts down the thread pool
func (tp *threadPool) Close() {
	tp.once.Do(func() {
		tp.lock.Lock()
		tp.closed = true
		tp.lock.Unlock()

		tp.cancelFunc()
		close(tp.tasks)
		tp.waitGroup.Wait()
		close(tp.errChan)
	})
}
