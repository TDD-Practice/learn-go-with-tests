package concurrency

import (
	"testing"
	"time"
)

func slowStubUrlChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckUrl(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "auto generated stub string"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckUrl(slowStubUrlChecker, urls)
	}
}
