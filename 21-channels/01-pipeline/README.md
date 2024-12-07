# Pipeline

With the help of goroutines and channels, construct a data processing pipeline. The pipeline consists of three stages: generator, adder, and collector. A function represents each stage. The specifications of the functions are the following.

The `generator(nums []int) <- chan int` function takes an integer array and pushes its elements into an integer channel.

The `adder(in <-chan int) <-chan float32` function increments digits by 3 of integers read from the `in` channel, and writes them to a `float32` channel.

The `collector(in <-chan float32) []float32` function reads numbers from the `in` channel and stores them in a `float32` array. When the channel closes, the function returns the array.

Insert the code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.

Hint: use [channels](https://go.dev/tour/concurrency/2).
