package channelbroadcaster

import "context"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func channelBroadcast(ctx context.Context, input <-chan any, outputs []chan<- any) {
	go func() {
		defer func() {
			for _, output := range outputs {
				close(output)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-input:
				if !ok {
					return
				}
				for _, output := range outputs {
					select {
					case output <- val:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()
}
