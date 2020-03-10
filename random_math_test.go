package base64Captcha

import (
	"bytes"
	"testing"
)

func TestRandText(t *testing.T) {
	type args struct {
		size        int
		sourceChars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{0, "foo"}, ""},
		{"", args{1, "aaa"}, "a"},
		{"", args{2, "bbb"}, "bb"},
		{"", args{3, "bbb"}, "bbb"},
		{"", args{3, "b"}, "bbb"},
		{"", args{4, "b"}, "bbbb"},
		{"", args{4, ""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandText(tt.args.size, tt.args.sourceChars); got != tt.want {
				t.Errorf("RandText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	type args struct {
		min int64
		max int64
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{-10, 10}},
		{"", args{-1, 5}},
		{"", args{0, 15}},
		{"", args{10, 14}},
		{"", args{10, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := random(tt.args.min, tt.args.max)
			// if out of bound then error
			if got < float64(tt.args.min) || got > float64(tt.args.max) {
				t.Errorf("RandText() = %v, out of range", got)
			}
		})
	}
}

func TestRandDarkAndLightColor(t *testing.T) {
	// Test RandColor
	for i := 0; i < 100; i++ {
		rgbA := RandColor()
		if rgbA.R < 0 || rgbA.R > 255 ||
			rgbA.B < 0 || rgbA.B > 255 ||
			rgbA.G < 0 || rgbA.G > 255 {
			t.Errorf("RandText() = %v, out of range", rgbA)
		}
	}

	// Test RandLightColor
	for i := 0; i < 100; i++ {
		rgbA := RandLightColor()
		if rgbA.R < 200 || rgbA.R > 255 ||
			rgbA.B < 200 || rgbA.B > 255 ||
			rgbA.G < 200 || rgbA.G > 255 {
			t.Errorf("RandText() = %v, out of range", rgbA)
		}
	}

	// Test RandDeepColor
	for i := 0; i < 100; i++ {
		rgbA := RandDeepColor()
		if rgbA.R < 0 || rgbA.R > 255 ||
			rgbA.B < 0 || rgbA.B > 255 ||
			rgbA.G < 0 || rgbA.G > 255 {
			t.Errorf("RandText() = %v, out of range", rgbA)
		}
	}
}

func TestRand(t *testing.T) {
	// test rand int Range
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{-10, 10}},
		{"", args{-1, 5}},
		{"", args{0, 15}},
		{"", args{10, 14}},
		{"", args{10, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randIntRange(tt.args.from, tt.args.to)

			// if out of bound then error
			if got < tt.args.from || got > tt.args.to {
				t.Errorf("RandText() = %v, out of range", got)
			}
		})
	}

	// test rand float Range

	type fargs struct {
		from float64
		to   float64
	}
	tests2 := []struct {
		name string
		arg  fargs
	}{
		{"", fargs{-10.0, 10.1}},
		{"", fargs{-1.0, 5.2}},
		{"", fargs{0, 15.3}},
		{"", fargs{10.1, 14.3}},
		{"", fargs{10.5, 10.5}},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := randFloat64Range(tt.arg.from, tt.arg.to)

			// if out of bound then error
			if got < tt.arg.from || got > tt.arg.to {
				t.Errorf("RandText() = %v, out of range", got)
			}
		})
	}
}

func TestRandomID(t *testing.T) {
	id := RandomId()
	if len(id) != idLen {
		t.Errorf("Wrong length got %d, want %d", len(id), idLen)
	}
	for _, val := range id {
		if !bytes.ContainsRune(idChars, val) {
			t.Errorf("got %v, want %v", idChars, val)
		}
	}
}

func TestRandBytes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{5}, 5},
		{"", args{0}, 0},
		{"", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randBytes(tt.args.n)

			// if out of bound then error
			if len(got) != tt.want {
				t.Errorf("randBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
