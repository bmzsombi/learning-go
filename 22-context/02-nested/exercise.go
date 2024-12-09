package subtask

import (
	"context"
	"time"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func SubTask(ctx context.Context) (result string, err error) {
	select {
	case <-ctx.Done(): // Context canceled before completing the task
		return "", ctx.Err()
	case <-time.After(200 * time.Millisecond): // Simulate long-running operation
		return "Subtask completed successfully", nil
	}
}

// StartTask starts the main task and handles its lifecycle.
func StartTask(ctx context.Context) (result string, err error) {
	// Create a new context with a 1-second timeout
	subCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel() // Ensure proper cleanup of the derived context

	resultChan := make(chan string) // Channel to receive results from SubTask
	errorChan := make(chan error)   // Channel to receive errors from SubTask

	// Start SubTask in a new goroutine
	go func() {
		defer close(resultChan)
		defer close(errorChan)
		res, err := SubTask(subCtx)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- res
		}
	}()

	select {
	case <-ctx.Done(): // Main context canceled
		return "", ctx.Err()
	case err := <-errorChan: // SubTask returned an error
		return "", err
	case res := <-resultChan: // SubTask completed successfully
		return "Main task status: " + res, nil
	}
}
