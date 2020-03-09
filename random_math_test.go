package base64Captcha

import (
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
