package base64Captcha

import (
	"image/color"
	"testing"

	"github.com/golang/freetype/truetype"
)

func TestDriverString_GenerateItem(t *testing.T) {
	type fields struct {
		Height          int
		Width           int
		NoiseTextCount  int
		NoiseDotCount   int
		ShowNoiseOption int
		CaptchaLen      int
		BgColor         *color.RGBA
		Fonts           []*truetype.Font
	}
	type args struct {
		content string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantItem Item
		wantErr  bool
	}{
		{"string", fields{80, 240, 20, 100, 2, 5, nil, fontsAll}, args{"45Ad8"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DriverString{
				Height:          tt.fields.Height,
				Width:           tt.fields.Width,
				NoiseCount:      tt.fields.NoiseTextCount,
				ShowLineOptions: tt.fields.ShowNoiseOption,
				Length:          tt.fields.CaptchaLen,
				BgColor:         tt.fields.BgColor,
				Fonts:           tt.fields.Fonts,
			}
			gotItem, err := d.GenerateItem(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("DriverString.GenerateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = ItemWriteToFile(gotItem, "_builds", tt.args.content, "png")
			if err != nil {
				t.Error(err)
			}

		})
	}
}
