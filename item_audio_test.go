// Copyright 2017 Eric Zhou. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base64Captcha

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func Test_newAudio(t *testing.T) {
	type args struct {
		id     string
		digits []byte
		lang   string
	}
	tests := []struct {
		name string
		args args
		want *ItemAudio
	}{
		{"zh3", args{RandomId(), randomDigits(3), "zh"}, nil},
		{"en4", args{RandomId(), randomDigits(4), "en"}, nil},
		{"ru2", args{RandomId(), randomDigits(2), "ru"}, nil},
		{"jp5", args{RandomId(), randomDigits(5), "jp"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newAudio(tt.args.id, tt.args.digits, tt.args.lang)
			if got == nil {
				t.Errorf("newAudio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemAudio_encodedLen(t *testing.T) {
	ia := newAudio(RandomId(), randomDigits(3), "zh")
	tests := []struct {
		name string
		a    *ItemAudio
		want int
	}{
		{"encode", ia, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.encodedLen(); got < tt.want {
				t.Errorf("ItemAudio.encodedLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemAudio_longestDigitSndLen(t *testing.T) {
	baseS := "0123456789abcdef"
	base := int64(len(baseS))
	num := time.Now().UnixNano()
	fmt.Printf("%x\n", num)
	newB := []byte{}
	for {
		idx := num % base
		bbb := []byte{byte(baseS[idx])}
		newB = append(bbb, newB...)
		num = num / base
		if num == 0 {
			break
		}
	}
	t.Log(string(newB))
}

func TestItemAudio_WriteTo(t *testing.T) {
	ia := newAudio(RandomId(), randomDigits(3), "zh")
	tests := []struct {
		name    string
		a       *ItemAudio
		wantN   int64
		wantW   string
		wantErr bool
	}{
		{"one", ia, 0, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			gotN, err := tt.a.WriteTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemAudio.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN < 1 {
				t.Errorf("ItemAudio.WriteTo() = %v, want %v", gotN, tt.wantN)
			}
			if gotW := w.String(); len(gotW) < 1 {
				t.Errorf("ItemAudio.WriteTo() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestItemAudio_EncodeB64string(t *testing.T) {
	ia := newAudio(RandomId(), randomDigits(5), "en")

	tests := []struct {
		name string
		a    *ItemAudio
		want string
	}{
		{"b64", ia, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.EncodeB64string(); len(got) < 1 {
				t.Errorf("ItemAudio.EncodeB64string() = %v, want %v", got, tt.want)
			}
		})
	}
}
