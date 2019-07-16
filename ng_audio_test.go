// Copyright 2011 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base64Captcha

import (
	"io/ioutil"
	"os"
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
