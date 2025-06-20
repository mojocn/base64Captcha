package rand

import (
	"math"
	"testing"
)

func TestFloat64(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Logf("got = %v", Float64())
	}
}

func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64()
	}
}

func TestIntn(t *testing.T) {
	t.Logf("got = %v", Intn(0))
	for i := 0; i < 20; i++ {
		t.Logf("got = %v", Intn(10))
	}
}

func TestInt31n(t *testing.T) {
	t.Logf("got = %v", Int31n(0))
	for i := 0; i < 20; i++ {
		t.Logf("got = %v", Int31n(10))
	}
}

func TestUint64n(t *testing.T) {
	t.Logf("got = %v", Uint64n(0))
	t.Logf("math.MaxUint64 = %v", uint64(math.MaxUint64))
	for i := 0; i < 20; i++ {
		t.Logf("got = %v", Uint64n(math.MaxUint64))
	}
}

func TestInt63n(t *testing.T) {
	t.Logf("got = %v", Int63n(0))
	for i := 0; i < 20; i++ {
		t.Logf("got = %v", Int63n(math.MaxInt64))
	}
}

func TestUint64(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Logf("got = %v", Uint64())
	}
}

func TestInt63(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Logf("got = %v", Int63())
	}
}
