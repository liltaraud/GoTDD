package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Simple incrementation test", func(t *testing.T) {
		c := NewCounter()
		c.Inc()
		c.Inc()
		c.Inc()

		assertCounter(t, c, 3)
	})

	t.Run("Concurrent run test", func(t *testing.T) {
		expectedCount := 1000
		c := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			go func(w *sync.WaitGroup) {
				c.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, c, expectedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("Expected %d, got %d", want, got.Value())
	}
}
