package base64Captcha

import (
	"image/color"
	"reflect"
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
				fontsStorage:    DefaultEmbeddedFonts,
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

func TestNewDriverMath(t *testing.T) {
	type args struct {
		height          int
		width           int
		noiseCount      int
		showLineOptions int
		bgColor         *color.RGBA
		fonts           []string
	}
	tests := []struct {
		name string
		args args
		want *DriverMath
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriverMath(tt.args.height, tt.args.width, tt.args.noiseCount, tt.args.showLineOptions, tt.args.bgColor, nil, tt.args.fonts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriverMath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverMath_ConvertFonts(t *testing.T) {
	tests := []struct {
		name string
		d    *DriverMath
		want *DriverMath
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ConvertFonts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DriverMath.ConvertFonts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverMath_GenerateIdQuestionAnswer(t *testing.T) {
	tests := []struct {
		name         string
		d            *DriverMath
		wantId       string
		wantQuestion string
		wantAnswer   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotQuestion, gotAnswer := tt.d.GenerateIdQuestionAnswer()
			if gotId != tt.wantId {
				t.Errorf("DriverMath.GenerateIdQuestionAnswer() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotQuestion != tt.wantQuestion {
				t.Errorf("DriverMath.GenerateIdQuestionAnswer() gotQuestion = %v, want %v", gotQuestion, tt.wantQuestion)
			}
			if gotAnswer != tt.wantAnswer {
				t.Errorf("DriverMath.GenerateIdQuestionAnswer() gotAnswer = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
