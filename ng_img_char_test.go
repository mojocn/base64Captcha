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
		configC.Mode = i % 5
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

func TestEngineCharUseSequencedCharacters(t *testing.T) {
	tc, _ := ioutil.TempDir("", "UseSequencedCharacters")
	defer os.Remove(tc)
	configSequencedCharacters := configC
	configSequencedCharacters.UseCJKFonts = true
	configSequencedCharacters.Mode = CaptchaModeUseSequencedCharacters
	configSequencedCharacters.SequencedCharacters = []string{
		"文件", "下载", "测试", "词组", "验证", "顺序",
	}
	configSequencedCharacters.CaptchaLen = 9
	im := EngineCharCreate(configSequencedCharacters)
	rep := im.Content
	for _, each := range configSequencedCharacters.SequencedCharacters {
		rep = strings.Replace(rep, each, "", -1)
	}
	if len([]rune(rep)) != 1 {
		t.Errorf("notgood: %v [rep=%v]", im.Content, rep)
	}
	fileName := strings.Trim(im.Content, "/+-+=?")
	err := CaptchaWriteToFile(im, tc, fileName, "png")
	if err != nil {
		t.Error(err)
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

	t.Run("TestCaptchaModeUseSequencedCharacters", func(t *testing.T) {
		configBad := configC
		configBad.CaptchaLen = 9
		configBad.Mode = CaptchaModeUseSequencedCharacters
		configBad.SequencedCharacters = nil
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
