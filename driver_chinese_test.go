package base64Captcha

import (
	"image/color"
	"reflect"
	"testing"
)

func TestNewDriverChinese(t *testing.T) {
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
		want *DriverChinese
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriverChinese(tt.args.height, tt.args.width, tt.args.noiseCount, tt.args.showLineOptions, tt.args.length, tt.args.source, tt.args.bgColor, nil, tt.args.fonts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriverChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverChinese_ConvertFonts(t *testing.T) {
	tests := []struct {
		name string
		d    *DriverChinese
		want *DriverChinese
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ConvertFonts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DriverChinese.ConvertFonts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverChinese_GenerateIdQuestionAnswer(t *testing.T) {
	tests := []struct {
		name        string
		d           *DriverChinese
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
				t.Errorf("DriverChinese.GenerateIdQuestionAnswer() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotContent != tt.wantContent {
				t.Errorf("DriverChinese.GenerateIdQuestionAnswer() gotContent = %v, want %v", gotContent, tt.wantContent)
			}
			if gotAnswer != tt.wantAnswer {
				t.Errorf("DriverChinese.GenerateIdQuestionAnswer() gotAnswer = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}

func TestDriverChinese_DrawCaptcha(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name     string
		d        *DriverChinese
		args     args
		wantItem Item
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, err := tt.d.DrawCaptcha(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("DriverChinese.DrawCaptcha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("DriverChinese.DrawCaptcha() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
