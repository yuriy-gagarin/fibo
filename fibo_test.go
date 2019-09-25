package fibo

import (
	"testing"
)

var low = []struct {
	inp uint
	out uint
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
	{20, 6765},
	{30, 832040},
}

var medium = []struct {
	inp uint
	out uint
}{
	{40, 102334155},
	{50, 12586269025},
}

var large = []struct {
	inp uint
	out uint
}{
	{40, 102334155},
	{50, 12586269025},
	{60, 1548008755920},
	{70, 190392490709135},
	{80, 23416728348467685},
	{90, 2880067194370816120},
}

// Recursive breaks really fast
func TestFiboRecursive(t *testing.T) {
	ins := low

	for _, c := range ins {
		if o := FiboRecursive(c.inp); o != c.out {
			t.Fatalf("expected %d, got %d", c.out, o)
		}
	}
}

// Always correct
func TestFiboIterative(t *testing.T) {
	ins := append(low[:0:0], low...)
	ins = append(ins, medium...)
	ins = append(ins, large...)

	for _, c := range ins {
		if o := FiboIterative(c.inp); o != c.out {
			t.Fatalf("expected %d, got %d", c.out, o)
		}
	}
}

// Binet and it's variations are innacurate on large Ns
func TestFiboBinet(t *testing.T) {
	ins := append(low[:0:0], low...)
	ins = append(ins, medium...)

	for _, c := range ins {
		if o := FiboBinet(c.inp); o != c.out {
			t.Fatalf("expected %d, got %d", c.out, o)
		}
	}
}

func TestFiboBinet2(t *testing.T) {
	ins := append(low[:0:0], low...)
	ins = append(ins, medium...)

	for _, c := range ins {
		if o := FiboBinet2(c.inp); o != c.out {
			t.Fatalf("expected %d, got %d", c.out, o)
		}
	}
}

func TestFiboBinet3(t *testing.T) {
	ins := append(low[:0:0], low...)
	ins = append(ins, medium...)

	for _, c := range ins {
		if o := FiboBinet3(c.inp); o != c.out {
			t.Fatalf("expected %d, got %d", c.out, o)
		}
	}
}

// Always correct
func TestFiboRec2(t *testing.T) {
	ins := append(low[:0:0], low...)
	ins = append(ins, medium...)
	ins = append(ins, large...)

	for _, c := range ins {
		if o := FiboRec2(c.inp); o != c.out {
			t.Fatalf("expected %d, got %d", c.out, o)
		}
	}
}
