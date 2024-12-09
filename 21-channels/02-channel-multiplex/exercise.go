package channelmultiplexer

import (
	"context"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func channelMultiplex(ctx context.Context, inputs []chan any) chan any {
	output := make(chan any) // Output channel

	go func() {
		defer close(output) // Ensure the output channel is closed when done

		for _, ch := range inputs {
			// Launch a goroutine to read from each input channel
			go func(c chan any) {
				for {
					select {
					case <-ctx.Done(): // Context cancellation
						return
					case val, ok := <-c: // Read from the input channel
						if !ok { // Input channel is closed
							return
						}
						select {
						case output <- val: // Send the value to the output channel
						case <-ctx.Done(): // Context cancellation
							return
						}
					}
				}
			}(ch)
		}

		// Wait for the context to be canceled to terminate all goroutines
		<-ctx.Done()
	}()

	return output
}
