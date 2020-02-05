package log_test

import (
	"strings"
	"testing"
)

func a(a int) int {
	if a > 0 {
		_ = 10 * 12
	}
	return b("Called from A")
}

func b(b string) int {
	b = strings.ToUpper(b)
	if len(b) > 1024 {
		b = strings.ToLower(b)
	}
	return len(b)
}
func BenchmarkNoTracingNoMetricsNoLogging(b *testing.B) {
	RunBenchmark(b, a)
}
