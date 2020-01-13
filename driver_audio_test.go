// Copyright 2017 Eric Zhou. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base64Captcha

import (
	"reflect"
	"testing"
)

func TestDriverAudio_DrawCaptcha(t *testing.T) {
	type fields struct {
		CaptchaLen int
		Language   string
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
		{"Audio", fields{4, "zh"}, args{"1234"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DriverAudio{
				Length:   tt.fields.CaptchaLen,
				Language: tt.fields.Language,
			}
			gotItem, err := d.DrawCaptcha(tt.args.content)
			if err != nil {
				t.Error(err)
			}
			itemWriteFile(gotItem, "_builds", tt.args.content, "wav")
		})
	}
}

func TestNewDriverAudio(t *testing.T) {
	type args struct {
		length   int
		language string
	}
	tests := []struct {
		name string
		args args
		want *DriverAudio
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriverAudio(tt.args.length, tt.args.language); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriverAudio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriverAudio_GenerateIdQuestionAnswer(t *testing.T) {
	tests := []struct {
		name   string
		d      *DriverAudio
		wantId string
		wantQ  string
		wantA  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotQ, gotA := tt.d.GenerateIdQuestionAnswer()
			if gotId != tt.wantId {
				t.Errorf("DriverAudio.GenerateIdQuestionAnswer() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotQ != tt.wantQ {
				t.Errorf("DriverAudio.GenerateIdQuestionAnswer() gotQ = %v, want %v", gotQ, tt.wantQ)
			}
			if gotA != tt.wantA {
				t.Errorf("DriverAudio.GenerateIdQuestionAnswer() gotA = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}
