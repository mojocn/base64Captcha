package base64Captcha

import (
	"reflect"
	"testing"
)

func Test_mixSound(t *testing.T) {
	type args struct {
		dst []byte
		src []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mixSound(tt.args.dst, tt.args.src)
		})
	}
}

func Test_setSoundLevel(t *testing.T) {
	type args struct {
		a     []byte
		level float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setSoundLevel(tt.args.a, tt.args.level)
		})
	}
}

func Test_changeSpeed(t *testing.T) {
	type args struct {
		a     []byte
		speed float64
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
			if got := changeSpeed(tt.args.a, tt.args.speed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("changeSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeSilence(t *testing.T) {
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
			if got := makeSilence(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSilence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversedSound(t *testing.T) {
	type args struct {
		a []byte
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
			if got := reversedSound(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversedSound() = %v, want %v", got, tt.want)
			}
		})
	}
}
