// Copyright 2011 Dmitry Chestnykh. All rights reserved.

// Use of this source code is governed by a MIT-style

// license that can be found in the LICENSE file.

package base64Captcha

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func BenchmarkNewAudio(b *testing.B) {
	b.StopTimer()
	d := randomDigits(DefaultLen)
	id := randomId()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		newAudio(id, d, "")
	}
}

func BenchmarkAudioWriteTo(b *testing.B) {
	b.StopTimer()
	d := randomDigits(DefaultLen)
	id := randomId()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		a := newAudio(id, d, "")
		n, _ := a.WriteTo(ioutil.Discard)
		b.SetBytes(n)
	}
}

func TestEngineAudioCreate(t *testing.T) {
	ta, _ := ioutil.TempDir("", "audio")
	defer os.RemoveAll(ta)

	//todo:: fix zh zero sound
	for i := 0; i < 100; i++ {
		idKey := randomId()
		au := EngineAudioCreate(idKey, configA)
		if err := CaptchaWriteToFile(au, ta, au.VerifyValue, "wav"); err != nil {
			t.Log(err)
		}
	}

}

func Test_newAudio(t *testing.T) {
	type args struct {
		id     string
		digits []byte
		lang   string
	}
	tests := []struct {
		name string
		args args
		want *Audio
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAudio(tt.args.id, tt.args.digits, tt.args.lang); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAudio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudio_encodedLen(t *testing.T) {
	tests := []struct {
		name string
		a    *Audio
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.encodedLen(); got != tt.want {
				t.Errorf("Audio.encodedLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudio_makeBackgroundSound(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		a    *Audio
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.makeBackgroundSound(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Audio.makeBackgroundSound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudio_randomizedDigitSound(t *testing.T) {
	type args struct {
		n byte
	}
	tests := []struct {
		name string
		a    *Audio
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.randomizedDigitSound(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Audio.randomizedDigitSound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudio_longestDigitSndLen(t *testing.T) {
	tests := []struct {
		name string
		a    *Audio
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.longestDigitSndLen(); got != tt.want {
				t.Errorf("Audio.longestDigitSndLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudio_randomSpeed(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		a    *Audio
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.randomSpeed(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Audio.randomSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudio_makeWhiteNoise(t *testing.T) {
	type args struct {
		length int
		level  uint8
	}
	tests := []struct {
		name string
		a    *Audio
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.makeWhiteNoise(tt.args.length, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Audio.makeWhiteNoise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mixSound(t *testing.T) {
	type args struct {
		dst []byte
		src []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mixSound(tt.args.dst, tt.args.src)
		})
	}
}

func Test_setSoundLevel(t *testing.T) {
	type args struct {
		a     []byte
		level float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setSoundLevel(tt.args.a, tt.args.level)
		})
	}
}

func Test_changeSpeed(t *testing.T) {
	type args struct {
		a     []byte
		speed float64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeSpeed(tt.args.a, tt.args.speed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("changeSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeSilence(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSilence(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSilence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversedSound(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversedSound(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversedSound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudio_WriteTo(t *testing.T) {
	tests := []struct {
		name    string
		a       *Audio
		wantN   int64
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			gotN, err := tt.a.WriteTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("Audio.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Audio.WriteTo() = %v, want %v", gotN, tt.wantN)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Audio.WriteTo() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestAudio_BinaryEncoding(t *testing.T) {
	tests := []struct {
		name string
		a    *Audio
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.BinaryEncoding(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Audio.BinaryEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
