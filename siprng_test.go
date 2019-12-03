// Copyright 2017 Eric Zhou. All Rights Reserved.

//

// Licensed under the Apache License, Version 2.0 (the "License");

// you may not use this file except in compliance with the License.

// You may obtain a copy of the License at

//

//      http://www.apache.org/licenses/LICENSE-2.0

//

// Unless required by applicable law or agreed to in writing, software

// distributed under the License is distributed on an "AS IS" BASIS,

// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

// See the License for the specific language governing permissions and

// limitations under the License.

package base64Captcha

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSiphash(t *testing.T) {
	good := uint64(0xe849e8bb6ffe2567)
	cur := siphash(0, 0, 0)
	if cur != good {
		t.Fatalf("siphash: expected %x, got %x", good, cur)
	}
}

func TestSiprng(t *testing.T) {
	m := make(map[uint64]interface{})
	var yes interface{}
	r := siprng{}
	r.Seed([16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	for i := 0; i < 100000; i++ {
		v := r.Uint64()
		if _, ok := m[v]; ok {
			t.Errorf("siphash: collision on %d: %x", i, v)
		}
		m[v] = yes
	}
}

func TestSiprngBytes(t *testing.T) {
	r := siprng{}
	r.Seed([16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	x := r.Bytes(32)
	if len(x) != 32 {
		t.Fatalf("siphash: wrong length: expected 32, got %d", len(x))
	}
	y := r.Bytes(32)
	if bytes.Equal(x, y) {
		t.Fatalf("siphash: stream repeats: %x = %x", x, y)
	}
	r.Seed([16]byte{})
	z := r.Bytes(32)
	if bytes.Equal(z, x) {
		t.Fatalf("siphash: outputs under different keys repeat: %x = %x", z, x)
	}
}

func BenchmarkSiprng(b *testing.B) {
	b.SetBytes(8)
	p := &siprng{}
	for i := 0; i < b.N; i++ {
		p.Uint64()
	}
}

func Test_siphash(t *testing.T) {
	type args struct {
		k0 uint64
		k1 uint64
		m  uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := siphash(tt.args.k0, tt.args.k1, tt.args.m); got != tt.want {
				t.Errorf("siphash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Seed(t *testing.T) {
	type args struct {
		k [16]byte
	}
	tests := []struct {
		name string
		p    *siprng
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Seed(tt.args.k)
		})
	}
}

func Test_siprng_Uint64(t *testing.T) {
	tests := []struct {
		name string
		p    *siprng
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Uint64(); got != tt.want {
				t.Errorf("siprng.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Bytes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		p    *siprng
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Bytes(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("siprng.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Int63(t *testing.T) {
	tests := []struct {
		name string
		p    *siprng
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Int63(); got != tt.want {
				t.Errorf("siprng.Int63() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Uint32(t *testing.T) {
	tests := []struct {
		name string
		p    *siprng
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Uint32(); got != tt.want {
				t.Errorf("siprng.Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Int31(t *testing.T) {
	tests := []struct {
		name string
		p    *siprng
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Int31(); got != tt.want {
				t.Errorf("siprng.Int31() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Intn(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		p    *siprng
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Intn(tt.args.n); got != tt.want {
				t.Errorf("siprng.Intn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Int63n(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		p    *siprng
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Int63n(tt.args.n); got != tt.want {
				t.Errorf("siprng.Int63n() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Int31n(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		p    *siprng
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Int31n(tt.args.n); got != tt.want {
				t.Errorf("siprng.Int31n() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Float64(t *testing.T) {
	tests := []struct {
		name string
		p    *siprng
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Float64(); got != tt.want {
				t.Errorf("siprng.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Int(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		p    *siprng
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Int(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("siprng.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_siprng_Float(t *testing.T) {
	type args struct {
		from float64
		to   float64
	}
	tests := []struct {
		name string
		p    *siprng
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Float(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("siprng.Float() = %v, want %v", got, tt.want)
			}
		})
	}
}
