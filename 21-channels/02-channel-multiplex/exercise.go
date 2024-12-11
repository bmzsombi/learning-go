package channelmultiplexer

import (
	"context"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func channelMultiplex(ctx context.Context, inputs []chan any) chan any {
	output := make(chan any)

	go func() {
		defer close(output)

		var done = make(chan struct{}, len(inputs))

		for _, ch := range inputs {
			ch := ch
			go func(ch chan any) {
				for {
					select {
					case <-ctx.Done():
						done <- struct{}{}
						return
					case val, ok := <-ch:
						if !ok {
							done <- struct{}{}
							return
						}
						select {
						case output <- val:
						case <-ctx.Done():
							done <- struct{}{}
							return
						}
					}
				}
			}(ch)
		}

		for range inputs {
			<-done
		}
	}()

	return output
}
