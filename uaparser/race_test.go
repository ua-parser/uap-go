package uaparser

import (
	"sync"
	"testing"
)

// TestNewConcurrent_NoRaceOnDefaultRegexDefinitions reproduces the data race
// reported in issue #100. Without the fix, multiple goroutines calling New()
// with the default cached RegexDefinitions racing on the shared []*uaParser
// pointers triggers a race detector failure under -race.
func TestNewConcurrent_NoRaceOnDefaultRegexDefinitions(t *testing.T) {
	const n = 16
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			if _, err := New(); err != nil {
				t.Errorf("New: %v", err)
			}
		}()
	}
	wg.Wait()
}
