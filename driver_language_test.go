package base64Captcha

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/golang/freetype/truetype"
)

func TestDriverLanguage_DrawCaptcha(t *testing.T) {
	ds := NewDriverLanguage(80, 240, 5, OptionShowSineLine|OptionShowSlimeLine|OptionShowHollowLine, 5, nil, nil, []*truetype.Font{fontChinese}, "emotion")

	for i := 0; i < 40; i++ {
		_, q, _ := ds.GenerateIdQuestionAnswer()
		item, err := ds.DrawCaptcha(q)
		if err != nil {
			t.Error(err)
		}
		itemWriteFile(item, "_builds", RandomId(), "png")
	}
}

func Test_generateRandomRune(t *testing.T) {
	type args struct {
		size int
		code string
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
			if got := generateRandomRune(tt.args.size, tt.args.code); got != tt.want {
				t.Errorf("generateRandomRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDriverLanguage(t *testing.T) {
	type args struct {
		height          int
		width           int
		noiseCount      int
		showLineOptions int
		length          int
		bgColor         *color.RGBA
		fonts           []*truetype.Font
		languageCode    string
	}
	tests := []struct {
		name string
		args args
		want *DriverLanguage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriverLanguage(tt.args.height, tt.args.width, tt.args.noiseCount, tt.args.showLineOptions, tt.args.length, tt.args.bgColor, nil, tt.args.fonts, tt.args.languageCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriverLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverLanguage_GenerateIdQuestionAnswer(t *testing.T) {
	tests := []struct {
		name        string
		d           *DriverLanguage
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
				t.Errorf("DriverLanguage.GenerateIdQuestionAnswer() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotContent != tt.wantContent {
				t.Errorf("DriverLanguage.GenerateIdQuestionAnswer() gotContent = %v, want %v", gotContent, tt.wantContent)
			}
			if gotAnswer != tt.wantAnswer {
				t.Errorf("DriverLanguage.GenerateIdQuestionAnswer() gotAnswer = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
