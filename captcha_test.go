package base64Captcha

import (
	"github.com/mojocn/base64Captcha/store"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

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
	testDir, _ := ioutil.TempDir("", "")
	defer os.Remove(testDir)

	for idx, vv := range []interface{}{configA, configD} {

		idkey, cap := GenerateCaptcha("", vv)
		ext := "png"
		if idx == 0 {
			ext = "wav"
		}

		CaptchaWriteToFile(cap, testDir, idkey, ext)
		CaptchaWriteToFile(cap, testDir, idkey, ext)

		CaptchaWriteToFile(cap, testDir, idkey, ext)

		// t.Log(idkey, globalStore.Get(idkey, false))

	}
	testDirAll, _ := ioutil.TempDir("", "all")
	defer os.RemoveAll(testDirAll)
	for i := 0; i < 16; i++ {
		configC.Mode = i % 4
		idkey, cap := GenerateCaptcha("", configC)
		ext := "png"
		err := CaptchaWriteToFile(cap, testDirAll, "char_"+idkey, ext)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestCaptchaWriteToBase64Encoding(t *testing.T) {
	_, cap := GenerateCaptcha("", configD)
	base64string := CaptchaWriteToBase64Encoding(cap)
	if !strings.Contains(base64string, MimeTypeCaptchaImage) {

		t.Error("encodeing base64 string failed.")
	}
	_, capA := GenerateCaptcha("", configA)
	base64stringA := CaptchaWriteToBase64Encoding(capA)
	if !strings.Contains(base64stringA, MimeTypeCaptchaAudio) {

		t.Error("encodeing base64 string failed.")
	}

}

func TestVerifyCaptcha(t *testing.T) {
	idkey, _ := GenerateCaptcha("", configD)
	verifyValue := globalStore.Get(idkey, false)
	if VerifyCaptcha(idkey, verifyValue) {
		t.Log(idkey, verifyValue)
	} else {
		t.Error("verify captcha content is failed.")
	}

	VerifyCaptcha("", "")
	VerifyCaptcha("dsafasf", "ddd")

}

func TestPathExists(t *testing.T) {

	testDir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(testDir)
	assert.True(t, pathExists(testDir))
	assert.False(t, pathExists(testDir+"/NotExistFolder"))
}

func TestCaptchaWriteToFileCreateDirectory(t *testing.T) {

	idKey, captcha := GenerateCaptcha("", configD)
	testDir, _ := ioutil.TempDir("", "")
	defer os.Remove(testDir)
	assert.Nil(t, CaptchaWriteToFile(captcha, testDir+"/NotExistFolder", idKey, "png"))
}

func TestCaptchaWriteToFileCreateFileFailed(t *testing.T) {

	var err error
	idKey, captcha := GenerateCaptcha("", configD)
	testDir, _ := ioutil.TempDir("", "")
	defer os.Remove(testDir)
	noPermissionDirPath := testDir + "/NoPermission"

	err = os.Mkdir(noPermissionDirPath, os.ModeDir)
	assert.Nil(t, err)

	err = CaptchaWriteToFile(captcha, noPermissionDirPath, idKey, "png")
	//has no permission must failed
	assert.Nil(t, err)
}

func TestSetCustomStore(t *testing.T) {
	s := store.NewMemoryStore(1000, 10*time.Minute)
	SetCustomStore(s)
	assert.Equal(t, globalStore, s)
}
