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
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(200 * time.Millisecond):
		return "Subtask completed successfully", nil
	}
}

// StartTask starts the main task and handles its lifecycle.
func StartTask(ctx context.Context) (result string, err error) {
	subCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	resultChan := make(chan string)
	errorChan := make(chan error)

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
	case <-ctx.Done():
		return "", ctx.Err()
	case err := <-errorChan:
		return "", err
	case res := <-resultChan:
		return "Main task status: " + res, nil
	}
}
