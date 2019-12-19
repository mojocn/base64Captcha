package base64Captcha

import (
	"bytes"
	"image"
	"image/color"
	"reflect"
	"testing"
)

func TestNewItemDigit(t *testing.T) {
	type args struct {
		width    int
		height   int
		dotCount int
		maxSkew  float64
	}
	tests := []struct {
		name string
		args args
		want *ItemDigit
	}{
		{"one", args{240, 80, 6, 0.8}, nil},
		{"one", args{240, 80, 5, 0.8}, nil},
		{"one", args{240, 80, 6, 0.8}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItemDigit(tt.args.width, tt.args.height, tt.args.dotCount, tt.args.maxSkew); got == nil {
				t.Errorf("NewItemDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemDigit_getRandomPalette(t *testing.T) {
	tests := []struct {
		name string
		m    *ItemDigit
		want color.Palette
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.getRandomPalette(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemDigit.getRandomPalette() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemDigit_calculateSizes(t *testing.T) {
	type args struct {
		width  int
		height int
		ncount int
	}
	tests := []struct {
		name string
		m    *ItemDigit
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

func TestItemDigit_drawHorizLine(t *testing.T) {
	type args struct {
		fromX    int
		toX      int
		y        int
		colorIdx uint8
	}
	tests := []struct {
		name string
		m    *ItemDigit
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

func TestItemDigit_drawCircle(t *testing.T) {
	type args struct {
		x        int
		y        int
		radius   int
		colorIdx uint8
	}
	tests := []struct {
		name string
		m    *ItemDigit
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

func TestItemDigit_fillWithCircles(t *testing.T) {
	type args struct {
		n         int
		maxradius int
	}
	tests := []struct {
		name string
		m    *ItemDigit
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

func TestItemDigit_strikeThrough(t *testing.T) {
	tests := []struct {
		name string
		m    *ItemDigit
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.strikeThrough()
		})
	}
}

func TestItemDigit_drawDigit(t *testing.T) {
	type args struct {
		digit []byte
		x     int
		y     int
	}
	tests := []struct {
		name string
		m    *ItemDigit
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

func TestItemDigit_distort(t *testing.T) {
	type args struct {
		amplude float64
		period  float64
	}
	tests := []struct {
		name string
		m    *ItemDigit
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

func TestItemDigit_randomBrightness(t *testing.T) {
	type args struct {
		c   color.RGBA
		max uint8
	}
	tests := []struct {
		name string
		m    *ItemDigit
		args args
		want color.RGBA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.randomBrightness(tt.args.c, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemDigit.randomBrightness() = %v, want %v", got, tt.want)
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

func TestItemDigit_EncodeBinary(t *testing.T) {

	idd := NewItemDigit(80, 300, 20, 0.25)
	idd.Paletted = image.NewPaletted(image.Rect(0, 0, 80, 300), idd.getRandomPalette())

	tests := []struct {
		name string
		m    *ItemDigit
		want []byte
	}{
		{"one", idd, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.EncodeBinary(); len(got) == 0 {
				t.Errorf("ItemDigit.EncodeBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemDigit_WriteTo(t *testing.T) {
	tests := []struct {
		name    string
		m       *ItemDigit
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
				t.Errorf("ItemDigit.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ItemDigit.WriteTo() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ItemDigit.WriteTo() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestItemDigit_EncodeB64string(t *testing.T) {
	tests := []struct {
		name string
		m    *ItemDigit
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.EncodeB64string(); got != tt.want {
				t.Errorf("ItemDigit.EncodeB64string() = %v, want %v", got, tt.want)
			}
		})
	}
}
