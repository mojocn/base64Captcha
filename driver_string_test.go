package base64Captcha

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/golang/freetype/truetype"
)

func TestDriverString_DrawCaptcha(t *testing.T) {
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
				fontsArray:      tt.fields.Fonts,
			}
			gotItem, err := d.DrawCaptcha(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("DriverString.DrawCaptcha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = itemWriteFile(gotItem, "_builds", tt.args.content, "png")
			if err != nil {
				t.Error(err)
			}

		})
	}
}

func TestNewDriverString(t *testing.T) {
	type args struct {
		height          int
		width           int
		noiseCount      int
		showLineOptions int
		length          int
		source          string
		bgColor         *color.RGBA
		fonts           []string
	}
	tests := []struct {
		name string
		args args
		want *DriverString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriverString(tt.args.height, tt.args.width, tt.args.noiseCount, tt.args.showLineOptions, tt.args.length, tt.args.source, tt.args.bgColor, nil, tt.args.fonts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriverString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverString_ConvertFonts(t *testing.T) {
	tests := []struct {
		name string
		d    *DriverString
		want *DriverString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ConvertFonts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DriverString.ConvertFonts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverString_GenerateIdQuestionAnswer(t *testing.T) {
	tests := []struct {
		name        string
		d           *DriverString
		wantId      string
		wantContent string
		wantAnswer  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotContent, gotAnswer := tt.d.GenerateIdQuestionAnswer()
			if gotId != tt.wantId {
				t.Errorf("DriverString.GenerateIdQuestionAnswer() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotContent != tt.wantContent {
				t.Errorf("DriverString.GenerateIdQuestionAnswer() gotContent = %v, want %v", gotContent, tt.wantContent)
			}
			if gotAnswer != tt.wantAnswer {
				t.Errorf("DriverString.GenerateIdQuestionAnswer() gotAnswer = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
