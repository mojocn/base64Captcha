package base64Captcha

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestEngineCharCreate(t *testing.T) {
	tc, _ := ioutil.TempDir("", "audio")
	defer os.Remove(tc)
	for i := 0; i < 16; i++ {
		configC.Mode = 5
		boooo := i%2 == 0
		configC.IsUseSimpleFont = boooo
		configC.IsShowSlimeLine = boooo
		configC.IsShowNoiseText = boooo
		configC.IsShowHollowLine = boooo
		configC.IsShowSineLine = boooo
		configC.IsShowNoiseDot = boooo

		if configC.Mode == CaptchaModeChinese {
			configC.UseCJKFonts = true
		} else {
			configC.UseCJKFonts = false
		}

		im := EngineCharCreate(configC)
		fileName := strings.Trim(im.Content, "/+-+=?")
		err := CaptchaWriteToFile(im, tc, fileName, "png")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestEngineCharCreateShowNoiseText(t *testing.T) {
	tc, _ := ioutil.TempDir("", "TestEngineCharCreateShowNoiseText")
	defer os.Remove(tc)
	configC.Mode = CaptchaModeNumber
	configC.IsShowNoiseText = true
	random_int := 0x123345
	for _, complex := range []int{CaptchaComplexLower, CaptchaComplexMedium, CaptchaComplexHigh, random_int} {
		configC.ComplexOfNoiseText = complex
		im := EngineCharCreate(configC)
		fileName := strings.Trim(im.Content, "/+-+=?")
		err := CaptchaWriteToFile(im, tc, fileName, "png")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestEngineCharCreatePanics(t *testing.T) {
	tc, _ := ioutil.TempDir("", "TestEngineCharCreateShowNoiseText")
	defer os.Remove(tc)

	t.Run("TestCaptchaLen_EQ_0", func(t *testing.T) {
		configBad := configC
		configBad.CaptchaLen = 0
		defer func() {
			if err := recover(); err != nil {
				//good panic
				return
			}
		}()
		_ = EngineCharCreate(configBad)
		t.Error("shell not be here")
	})
	t.Run("CaptchaModeChinese without UseCJKFonts", func(t *testing.T) {
		configBad := configC
		configBad.Mode = CaptchaModeChinese
		configBad.UseCJKFonts = false
		defer func() {
			if err := recover(); err != nil {
				//good panic
				return
			}
		}()
		_ = EngineCharCreate(configBad)
		t.Error("shell not be here")
	})
}

func TestMath(t *testing.T) {
	for i := 0; i < 100; i++ {
		q, r := randArithmetic()
		t.Log(q, "--->", r)
	}
}

func TestEngineCharCreateStrList(t *testing.T) {
	tc, _ := ioutil.TempDir("", "audio")
	defer os.Remove(tc)

	configC.Mode = CaptchaModeUseRunePairs
	configC.UseCJKFonts = true
	configC.CaptchaLen = 9
	im := EngineCharCreate(configC)
	fileName := strings.Trim(im.Content, "/+-+=?")
	err := CaptchaWriteToFile(im, tc, fileName, "png")
	if err != nil {
		t.Error(err)
	}
}
