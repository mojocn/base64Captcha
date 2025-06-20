package rand

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
)

func Float64() float64 {
	var buf [8]byte
	_, _ = rand.Read(buf[:])

	// Convert bytes to uint64
	u := binary.BigEndian.Uint64(buf[:])

	// Only use the lower 53 bits (float64 mantissa precision)
	const max53bit = 1 << 53
	v := u >> (64 - 53) // keep top 53 bits

	return float64(v) / float64(max53bit)
}

func Intn(n int) int {
	if n <= 0 {
		return 0
	}

	// Use math/big to generate a secure random number less than n
	max := big.NewInt(int64(n))
	r, _ := rand.Int(rand.Reader, max)
	return int(r.Int64())
}

func Int31n(n int32) int32 {
	if n <= 0 {
		return 0
	}

	// Use math/big to generate a secure random number less than n
	max := big.NewInt(int64(n))
	r, _ := rand.Int(rand.Reader, max)
	return int32(r.Int64())
}

func Int63n(n int64) int64 {
	if n <= 0 {
		return 0
	}

	max := big.NewInt(int64(n))
	r, _ := rand.Int(rand.Reader, max)
	return int64(r.Int64())
}

func Int63() int64 {
	return Int63n(math.MaxInt64)
}

func Uint64n(n uint64) uint64 {
	if n == 0 {
		return 0
	}

	max := new(big.Int).SetUint64(n)
	r, _ := rand.Int(rand.Reader, max)
	return r.Uint64()
}

func Uint64() uint64 {
	return Uint64n(math.MaxUint64)
}
