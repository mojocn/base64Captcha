package base64Captcha

import (
	"io/ioutil"
	"testing"
)

func TestEngineDigitsCreate(t *testing.T) {
	td, _ := ioutil.TempDir("", "audio")

	for i := 0; i < 14; i++ {
		idKey := randomId()
		im := EngineDigitsCreate(idKey, configD)
		err := CaptchaWriteToFile(im, td, idKey, "png")
		if err != nil {
			t.Error(err)
		}
	}
}
