package fibo

import "testing"

func BenchmarkFiboRecursive_smallN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboRecursive(7)
	}
}

func BenchmarkFiboRecursive_largeN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboRecursive(20)
	}
}

func BenchmarkFiboIterative_smallN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboIterative(7)
	}
}

func BenchmarkFiboIterative_largeN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboIterative(20)
	}
}

func BenchmarkFiboBinet_smallN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboBinet(7)
	}
}

func BenchmarkFiboBinet_largeN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboBinet(20)
	}
}

func BenchmarkFiboBinet2_smallN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboBinet2(7)
	}
}

func BenchmarkFiboBinet2_largeN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboBinet2(20)
	}
}

func BenchmarkFiboBinet3_smallN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboBinet3(7)
	}
}

func BenchmarkFiboBinet3_largeN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboBinet3(20)
	}
}

func BenchmarkFiboRec2_smallN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboRec2(7)
	}
}

func BenchmarkFiboRec2_largeN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboRec2(20)
	}
}
