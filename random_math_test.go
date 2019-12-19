package base64Captcha

import (
	"image/color"
	"reflect"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandText(tt.args.size, tt.args.sourceChars); got != tt.want {
				t.Errorf("RandText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_random(t *testing.T) {
	type args struct {
		min int64
		max int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := random(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("random() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandDeepColor(t *testing.T) {
	tests := []struct {
		name string
		want color.RGBA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandDeepColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandDeepColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandLightColor(t *testing.T) {
	tests := []struct {
		name string
		want color.RGBA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandLightColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandLightColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandColor(t *testing.T) {
	tests := []struct {
		name string
		want color.RGBA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
