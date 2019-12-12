// Copyright 2017 Eric Zhou. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base64Captcha

import (
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
