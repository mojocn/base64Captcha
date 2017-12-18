package base64Captcha

import (
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

		t.Log(idkey, globalStore.Get(idkey, false))

	}

	for i := 0; i < 16; i++ {
		configC.Mode = i % 4
		idkey, cap := GenerateCaptcha("", configC)
		ext := "png"
		CaptchaWriteToFile(cap, GoTestOutputDir+"/all", "char_"+idkey, ext)

		t.Log(idkey, globalStore.Get(idkey, false))

	}
}
