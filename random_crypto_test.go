package base64Captcha

import (
	"reflect"
	"testing"
)

func TestRandomId(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(randomDigits(5))
	}
}
func TestRandomDigits(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(randomDigits(5))
	}
}
func TestParseDigitsToString(t *testing.T) {
	for i := 0; i < 10; i++ {
		byss := randomDigits(5)
		t.Log(byss)
		bsssstring := parseDigitsToString(byss)
		t.Log(bsssstring)
	}
}

func Test_deriveSeed(t *testing.T) {
	type args struct {
		purpose byte
		id      string
		digits  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantOut [16]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := deriveSeed(tt.args.purpose, tt.args.id, tt.args.digits); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("deriveSeed() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_randomDigits(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomDigits(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomBytes(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name  string
		args  args
		wantB []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := randomBytes(tt.args.length); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("randomBytes() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func Test_randomBytesMod(t *testing.T) {
	type args struct {
		length int
		mod    byte
	}
	tests := []struct {
		name  string
		args  args
		wantB []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := randomBytesMod(tt.args.length, tt.args.mod); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("randomBytesMod() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func Test_randomId(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomId(); got != tt.want {
				t.Errorf("randomId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseDigitsToString(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDigitsToString(tt.args.bytes); got != tt.want {
				t.Errorf("parseDigitsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
