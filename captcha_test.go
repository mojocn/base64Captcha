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
	"reflect"
	"testing"
)

func TestCaptcha_GenerateB64s(t *testing.T) {
	type fields struct {
		Driver Driver
		Store  Store
	}

	dDigit := DriverDigit{80, 240, 5, 0.7, 5}
	audioDriver := NewDriverAudio(rand.Intn(5), "en")
	tests := []struct {
		name     string
		fields   fields
		wantId   string
		wantB64s string
		wantErr  bool
	}{
		{"mem-digit", fields{&dDigit, DefaultMemStore}, "xxxx", "", false},
		{"mem-audio", fields{audioDriver, DefaultMemStore}, "xxxx", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCaptcha(tt.fields.Driver, tt.fields.Store)
			gotId, b64s, _, err := c.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Captcha.Generate() error = %v, wantErr %v", err, tt.wantErr)
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

func TestNewCaptcha(t *testing.T) {
	type args struct {
		driver Driver
		store  Store
	}
	tests := []struct {
		name string
		args args
		want *Captcha
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCaptcha(tt.args.driver, tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptcha_Generate(t *testing.T) {
	tests := []struct {
		name     string
		c        *Captcha
		wantId   string
		wantB64s string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotB64s, _, err := tt.c.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Captcha.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Captcha.Generate() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotB64s != tt.wantB64s {
				t.Errorf("Captcha.Generate() gotB64s = %v, want %v", gotB64s, tt.wantB64s)
			}
		})
	}
}
