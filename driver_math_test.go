package base64Captcha

import (
	"image/color"
	"testing"

	"github.com/golang/freetype/truetype"
)

func TestDriverMath_GenerateItem(t *testing.T) {
	type fields struct {
		Height          int
		Width           int
		NoiseCount      int
		ShowLineOptions int
		BgColor         *color.RGBA
		Fonts           []*truetype.Font
	}
	type args struct {
		question string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantItem Item
		wantErr  bool
	}{
		{"Math",
			fields{80, 240, 5, OptionShowSineLine | OptionShowSlimeLine | OptionShowHollowLine, nil, fontsAll},
			args{""},
			nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DriverMath{
				Height:          tt.fields.Height,
				Width:           tt.fields.Width,
				NoiseCount:      tt.fields.NoiseCount,
				ShowLineOptions: tt.fields.ShowLineOptions,
				BgColor:         tt.fields.BgColor,
				Fonts:           tt.fields.Fonts,
			}
			q, a := d.GenerateQuestionAnswer()
			gotItem, err := d.GenerateItem(q)
			if (err != nil) != tt.wantErr {
				t.Errorf("DriverMath.GenerateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			itemWriteFile(gotItem, "_builds", a, "png")

		})
	}
}
