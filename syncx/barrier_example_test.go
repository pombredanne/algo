package syncx_test

import (
	"fmt"
	"github.com/larsmans/algo/syncx"
)

func ExampleBarrier() {
	nworkers, nrounds := 3, 2
	bar := syncx.NewBarrier(nworkers + 1)

	for i := 0; i < nworkers; i++ {
		go func() {
			for j := 1; j <= nrounds; j++ {
				// Maybe do some work here.
				fmt.Println("Finished round", j)
				bar.Wait()
			}
		}()
	}

	// Waiting for bar here saves allocating an additional WaitGroup.
	for j := 1; j <= nrounds; j++ {
		bar.Wait()
	}
	fmt.Printf("All %d workers done", bar.Num()-1)

	// Output:
	// Finished round 1
	// Finished round 1
	// Finished round 1
	// Finished round 2
	// Finished round 2
	// Finished round 2
	// All 3 workers done
}
