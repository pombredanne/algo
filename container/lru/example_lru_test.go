package lru_test

import (
	"fmt"
	"github.com/larsmans/algo/container/lru"
	"time"
)

// Simulates loading a file from a slow disk.
func loadFile(name interface{}) (contents interface{}) {
	namestr := name.(string)
	fmt.Printf("Loading %q\n", namestr)
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("contents of %q", namestr)
}

func Example_lru() {
	cache := lru.New(loadFile, 3)
	for _, name := range []string{"hello.txt", "foo.txt", "bar.txt",
		"hello.txt", "bar.txt", "goodbye.txt"} {
		contents := cache.Get(name).(string)
		fmt.Printf("Got %s\n", contents)
	}
	// Output:
	// Loading "hello.txt"
	// Got contents of "hello.txt"
	// Loading "foo.txt"
	// Got contents of "foo.txt"
	// Loading "bar.txt"
	// Got contents of "bar.txt"
	// Got contents of "hello.txt"
	// Got contents of "bar.txt"
	// Loading "goodbye.txt"
	// Got contents of "goodbye.txt"
}
