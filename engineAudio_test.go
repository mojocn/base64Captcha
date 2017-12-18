// Copyright 2011 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base64Captcha

import (
	"io/ioutil"
	"testing"
)

func BenchmarkNewAudio(b *testing.B) {
	b.StopTimer()
	d := RandomDigits(DefaultLen)
	id := RandomId()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewAudio(id, d, "")
	}
}

func BenchmarkAudioWriteTo(b *testing.B) {
	b.StopTimer()
	d := RandomDigits(DefaultLen)
	id := RandomId()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		a := NewAudio(id, d, "")
		n, _ := a.WriteTo(ioutil.Discard)
		b.SetBytes(n)
	}
}

func TestEngineAudioCreate(t *testing.T) {
	for i := 0; i < 10; i++ {
		idKey := RandomId()
		au := EngineAudioCreate(idKey, configA)
		if err := CaptchaWriteToFile(au, GoTestOutputDir+"/audio", idKey, "wav"); err != nil {
			t.Log(err)
		}
	}

}
