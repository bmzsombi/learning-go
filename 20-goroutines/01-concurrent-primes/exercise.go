package concurrentprimes

import "time"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func GeneratePrimes(n int) []int {
	if n < 2 {
		return []int{}
	}

	numChan := make(chan int)
	primesChan := make(chan int)

	go func() {
		defer close(numChan)
		for i := 2; i <= n; i++ {
			numChan <- i
		}
	}()

	go func() {
		defer close(primesChan)
		primesChan <- 2
		primes := []int{2}

		for num := range numChan {
			isPrime := true
			for _, p := range primes {
				if num%p == 0 {
					isPrime = false
					break
				}
				if p*p > num {
					break
				}
			}
			if isPrime {
				primes = append(primes, num)
				primesChan <- num
			}
		}
	}()

	primes := []int{}
	for prime := range primesChan {
		primes = append(primes, prime)
	}

	time.Sleep(50 * time.Millisecond)
	return primes
}
