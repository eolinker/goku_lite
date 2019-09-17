package internal

import (
	"testing"
	"time"
)

func TestRetryBackoff(t *testing.T) {
	RegisterTestingT(t)

	for i := -1; i <= 16; i++ {
		backoff := RetryBackoff(i, time.Millisecond, 512*time.Millisecond)
		Expect(backoff >= 0).To(BeTrue())
		Expect(backoff <= 512*time.Millisecond).To(BeTrue())
	}
}
