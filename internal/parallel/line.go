package parallel

import (
	"runtime"
	"sync"
)

// ProcessInParallel dispatches a parameter fn into multiple goroutines by splitting the parameter length
// by the number of available CPUs and assigning the length parts into each fn.
func ProcessInParallel(length int, fn func(start, end int)) {
	processors := runtime.GOMAXPROCS(0)
	counter := length
	partSize := length / processors

	if processors <= 1 || partSize <= processors {
		fn(0, length)
	} else {
		var wg sync.WaitGroup
		for counter > 0 {
			start := counter - partSize
			end := counter
			if start < 0 {
				start = 0
			}

			counter -= partSize
			wg.Add(1)

			go func(start, end int) {
				defer wg.Done()
				fn(start, end)
			}(start, end)
		}

		wg.Wait()
	}
}
