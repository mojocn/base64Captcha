package base64Captcha

import (
	"testing"
)

func TestEngineCharCreate(t *testing.T) {

	for i := 0; i < 16; i++ {
		configC.Mode = i % 4
		boooo := i%2 == 0
		configC.IsUseSimpleFont = boooo
		configC.IsShowSlimeLine = boooo
		configC.IsShowNoiseText = boooo
		configC.IsShowHollowLine = boooo
		configC.IsShowSineLine = boooo
		configC.IsShowNoiseDot = boooo

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
