// Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions.

package popcounttest

import (
	"testing"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func expressionPopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func loopPopCount(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(8*i))])
	}
	return count
}

func BenchmarkExpressionPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expressionPopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loopPopCount(0x1234567890ABCDEF)
	}
}
