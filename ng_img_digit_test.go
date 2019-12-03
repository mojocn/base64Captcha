package base64Captcha

import (
	"bytes"
	"image/color"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestEngineDigitsCreate(t *testing.T) {
	td, _ := ioutil.TempDir("", "audio")
	defer os.Remove(td)
	for i := 0; i < 14; i++ {
		idKey := randomId()
		im := EngineDigitsCreate(idKey, configD)
		err := CaptchaWriteToFile(im, td, idKey, "png")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestCaptchaImageDigit_getRandomPalette(t *testing.T) {
	tests := []struct {
		name string
		m    *CaptchaImageDigit
		want color.Palette
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.getRandomPalette(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptchaImageDigit.getRandomPalette() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptchaImageDigit_calculateSizes(t *testing.T) {
	type args struct {
		width  int
		height int
		ncount int
	}
	tests := []struct {
		name string
		m    *CaptchaImageDigit
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.calculateSizes(tt.args.width, tt.args.height, tt.args.ncount)
		})
	}
}

func TestCaptchaImageDigit_drawHorizLine(t *testing.T) {
	type args struct {
		fromX    int
		toX      int
		y        int
		colorIdx uint8
	}
	tests := []struct {
		name string
		m    *CaptchaImageDigit
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.drawHorizLine(tt.args.fromX, tt.args.toX, tt.args.y, tt.args.colorIdx)
		})
	}
}

func TestCaptchaImageDigit_drawCircle(t *testing.T) {
	type args struct {
		x        int
		y        int
		radius   int
		colorIdx uint8
	}
	tests := []struct {
		name string
		m    *CaptchaImageDigit
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.drawCircle(tt.args.x, tt.args.y, tt.args.radius, tt.args.colorIdx)
		})
	}
}

func TestCaptchaImageDigit_fillWithCircles(t *testing.T) {
	type args struct {
		n         int
		maxradius int
	}
	tests := []struct {
		name string
		m    *CaptchaImageDigit
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.fillWithCircles(tt.args.n, tt.args.maxradius)
		})
	}
}

func TestCaptchaImageDigit_strikeThrough(t *testing.T) {
	tests := []struct {
		name string
		m    *CaptchaImageDigit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.strikeThrough()
		})
	}
}

func TestCaptchaImageDigit_drawDigit(t *testing.T) {
	type args struct {
		digit []byte
		x     int
		y     int
	}
	tests := []struct {
		name string
		m    *CaptchaImageDigit
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.drawDigit(tt.args.digit, tt.args.x, tt.args.y)
		})
	}
}

func TestCaptchaImageDigit_distort(t *testing.T) {
	type args struct {
		amplude float64
		period  float64
	}
	tests := []struct {
		name string
		m    *CaptchaImageDigit
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.distort(tt.args.amplude, tt.args.period)
		})
	}
}

func TestCaptchaImageDigit_randomBrightness(t *testing.T) {
	type args struct {
		c   color.RGBA
		max uint8
	}
	tests := []struct {
		name string
		m    *CaptchaImageDigit
		args args
		want color.RGBA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.randomBrightness(tt.args.c, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptchaImageDigit.randomBrightness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min3(t *testing.T) {
	type args struct {
		x uint8
		y uint8
		z uint8
	}
	tests := []struct {
		name  string
		args  args
		wantM uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := min3(tt.args.x, tt.args.y, tt.args.z); gotM != tt.wantM {
				t.Errorf("min3() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func Test_max3(t *testing.T) {
	type args struct {
		x uint8
		y uint8
		z uint8
	}
	tests := []struct {
		name  string
		args  args
		wantM uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := max3(tt.args.x, tt.args.y, tt.args.z); gotM != tt.wantM {
				t.Errorf("max3() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestCaptchaImageDigit_BinaryEncoding(t *testing.T) {
	tests := []struct {
		name string
		m    *CaptchaImageDigit
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.BinaryEncoding(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptchaImageDigit.BinaryEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptchaImageDigit_WriteTo(t *testing.T) {
	tests := []struct {
		name    string
		m       *CaptchaImageDigit
		want    int64
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			got, err := tt.m.WriteTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("CaptchaImageDigit.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CaptchaImageDigit.WriteTo() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("CaptchaImageDigit.WriteTo() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
