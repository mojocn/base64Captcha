// Copyright 2017 Eric Zhou. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package base64Captcha supports digits, numbers,alphabet, arithmetic, audio and digit-alphabet captcha.
// base64Captcha is used for fast development of RESTful APIs, web apps and backend services in Go. give a string identifier to the package and it returns with a base64-encoding-png-string
package base64Captcha

import (
	"math/rand"
	"testing"
)

func TestCaptcha_GenerateB64s(t *testing.T) {
	type fields struct {
		Driver Driver
		Store  Store
	}

	dDigit := driverDigit{80, 240, 5, 0.7, 5}
	audioDriver := NewDriverAudio(rand.Intn(5), "en")
	driverChar := NewDriverString(80, 240, 4, OptionShowHollowLine|OptionShowSlimeLine|OptionShowSineLine, 4, nil, fontsAll)
	driverChinese := NewDriverLanguage(*driverChar,"zh-CN")
	tests := []struct {
		name     string
		fields   fields
		wantId   string
		wantB64s string
		wantErr  bool
	}{
		{"mem-digit", fields{&dDigit, DefaultMemStore}, "xxxx", "", false},
		{"mem-audio", fields{audioDriver, DefaultMemStore}, "xxxx", "", false},
		{"mem-char", fields{driverChar, DefaultMemStore}, "xxxx", "", false},
		{"mem-chinese", fields{driverChinese, DefaultMemStore}, "xxxx", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCaptcha(tt.fields.Driver, tt.fields.Store)
			gotId, b64s, err := c.GenerateB64s()
			if (err != nil) != tt.wantErr {
				t.Errorf("Captcha.GenerateB64s() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(b64s)

			a := c.Store.Get(gotId, false)
			if !c.Verify(gotId, a, true) {
				t.Error("false")
			}
		})
	}
}

func TestCaptcha_Verify(t *testing.T) {
	type fields struct {
		Driver Driver
		Store  Store
	}
	type args struct {
		id     string
		answer string
		clear  bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantMatch bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Captcha{
				Driver: tt.fields.Driver,
				Store:  tt.fields.Store,
			}
			if gotMatch := c.Verify(tt.args.id, tt.args.answer, tt.args.clear); gotMatch != tt.wantMatch {
				t.Errorf("Captcha.Verify() = %v, want %v", gotMatch, tt.wantMatch)
			}
		})
	}
}
