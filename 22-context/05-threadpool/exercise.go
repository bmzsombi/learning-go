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

type threadPool struct {
	workerLimit chan struct{}   // Szálkorlátozó
	tasks       chan Runnable   // Feladatok csatornája
	wg          sync.WaitGroup  // Szinkronizáció
	closeOnce   sync.Once       // Egyszeri lezárás biztosítása
	errChan     chan error      // Hibák továbbítása
	ctx         context.Context // Kontextus a leállításhoz
	cancel      context.CancelFunc
}

func NewThreadPool(n int) (threadPool, chan error) {
	ctx, cancel := context.WithCancel(context.Background())
	pool := &threadPool{
		workerLimit: make(chan struct{}, n),
		tasks:       make(chan Runnable),
		errChan:     make(chan error, 100),
		ctx:         ctx,
		cancel:      cancel,
	}

	// Indítsuk el a dolgozókat
	for i := 0; i < n; i++ {
		go pool.worker()
	}

	return pool, pool.errChan
}

func (tp *threadPool) worker() {
	for {
		select {
		case <-tp.ctx.Done():
			// Kilépés, ha a pool leáll
			return
		case task, ok := <-tp.tasks:
			if !ok {
				// A feladatcsatorna bezárt
				return
			}

			// Indítsuk a feladatot
			tp.workerLimit <- struct{}{} // Szál foglalása
			tp.wg.Add(1)

			go func(task Runnable) {
				defer func() {
					<-tp.workerLimit // Szál felszabadítása
					tp.wg.Done()
				}()

				if err := task.Run(tp.ctx); err != nil {
					// Hiba továbbítása, ha a csatorna megtelt, logoljunk
					select {
					case tp.errChan <- err:
					default:
						log.Println("Error channel full, dropping error:", err)
					}
				}
			}(task)
		}
	}
}

func (tp *threadPool) Run(task Runnable) {
	select {
	case <-tp.ctx.Done():
		// Ha a pool már le van zárva
		return
	case tp.tasks <- task:
		// Feladat hozzáadva
	}
}

func (tp *threadPool) Close() {
	tp.closeOnce.Do(func() {
		tp.cancel()       // Kontextus lezárása
		close(tp.tasks)   // Feladatok csatornájának bezárása
		tp.wg.Wait()      // Várakozás az összes szál befejezésére
		close(tp.errChan) // Hibacsatorna lezárása
	})
}
