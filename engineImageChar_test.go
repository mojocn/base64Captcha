package base64Captcha

import (
	"testing"
)

func TestEngineCharCreate(t *testing.T) {

	for i := 0; i < 16; i++ {
		configC.Mode = i % 4
		im := EngineCharCreate(configC)
		CaptchaWriteToFile(im, GoTestOutputDir+"/char", im.Content, "png")
		t.Log(im.Content, im.VerifyValue)
	}
}
func TestMath(t *testing.T) {
	for i := 0; i < 100; i++ {
		q, r := randArithmetic()
		t.Log(q, "--->", r)
	}
}
