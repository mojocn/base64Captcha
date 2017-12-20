package base64Captcha

import (
	"testing"
)

func TestEngineDigitsCreate(t *testing.T) {

	for i := 0; i < 14; i++ {
		idKey := randomId()
		im := EngineDigitsCreate(idKey, configD)
		CaptchaWriteToFile(im, GoTestOutputDir+"/digit", idKey, "png")
		t.Log(idKey, im.VerifyValue)
	}
}
