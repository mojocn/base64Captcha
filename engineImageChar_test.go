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
		configC.Mode = i % 4
		boooo := i%2 == 0
		configC.IsUseSimpleFont = boooo
		configC.IsShowSlimeLine = boooo
		configC.IsShowNoiseText = boooo
		configC.IsShowHollowLine = boooo
		configC.IsShowSineLine = boooo
		configC.IsShowNoiseDot = boooo

		im := EngineCharCreate(configC)
		fileName := strings.Trim(im.Content, "/+-+=?")
		err := CaptchaWriteToFile(im, tc, fileName, "png")
		if err != nil {
			t.Error(err)
		}
	}
}
func TestMath(t *testing.T) {
	for i := 0; i < 100; i++ {
		q, r := randArithmetic()
		t.Log(q, "--->", r)
	}
}
