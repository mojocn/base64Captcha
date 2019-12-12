package base64Captcha

import (
	"image/color"
	"testing"
)

func TestDriverMath_DrawCaptcha(t *testing.T) {
	type fields struct {
		Height          int
		Width           int
		NoiseCount      int
		ShowLineOptions int
		BgColor         *color.RGBA
		Fonts           []string
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
			fields{80, 240, 5, OptionShowSineLine | OptionShowSlimeLine | OptionShowHollowLine, nil, []string{"3Dumb.ttf"}},
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
			d.ConvertFonts()
			_, q, a := d.GenerateIdQuestionAnswer()

			gotItem, err := d.DrawCaptcha(q)
			if (err != nil) != tt.wantErr {
				t.Errorf("DriverMath.DrawCaptcha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			itemWriteFile(gotItem, "_builds", a, "png")

		})
	}
}
