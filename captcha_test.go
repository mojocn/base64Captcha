package base64Captcha

import (
	"strings"
	"testing"
)

const ExampleFontDirPath = "/Users/ericzhou/go/src/github.com/mojocn/base64Captcha/examples/fonts"
const ExampleFontExtension = "ttf"

var configD = ConfigDigit{
	Height:     80,
	Width:      240,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 5,
}

var configA = ConfigAudio{
	CaptchaLen: 6,
	Language:   "zh",
}

var configC = ConfigCharacter{
	Height:             60,
	Width:              240,
	Mode:               0,
	ComplexOfNoiseText: 0,
	ComplexOfNoiseDot:  0,
	IsUseSimpleFont:    false,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         6,
}

func TestGenerateCaptcha(t *testing.T) {
	for idx, vv := range []interface{}{configA, configD} {

		idkey, cap := GenerateCaptcha("", vv)
		ext := "png"
		if idx == 0 {
			ext = "wav"
		}
		CaptchaWriteToFile(cap, GoTestOutputDir+"/all", idkey, ext)

		//t.Log(idkey, globalStore.Get(idkey, false))

	}

	for i := 0; i < 16; i++ {
		configC.Mode = i % 4
		idkey, cap := GenerateCaptcha("", configC)
		ext := "png"
		err := CaptchaWriteToFile(cap, GoTestOutputDir+"/all", "char_"+idkey, ext)
		if err == nil {
			t.Log(idkey, globalStore.Get(idkey, false))
		} else {
			t.Error(idkey)
		}
	}
}

func TestCaptchaWriteToBase64Encoding(t *testing.T) {
	idkey, cap := GenerateCaptcha("", configD)
	base64string := CaptchaWriteToBase64Encoding(cap)
	if strings.Contains(base64string, "base64,") {
		t.Log(base64string, idkey)
	} else {
		t.Error("encodeing base64 string failed.")
	}

}

func TestVerifyCaptcha(t *testing.T) {
	idkey, _ := GenerateCaptcha("", configD)
	verifyValue := globalStore.Get(idkey, false)
	if verifyValue == "" {
		t.Error("verify captcha content is failed.")
	} else {
		t.Log(verifyValue)
	}
}
