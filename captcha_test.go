package base64Captcha

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha/store"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"runtime"
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

func TestGenerateCaptchaOnFailed(t *testing.T) {

	preStore := globalStore
	s := store.NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6378",
	}), "test:", 5*time.Minute)
	SetCustomStore(s)
	idkey, capt := GenerateCaptcha("", configA)
	assert.Equal(t, "", idkey)
	assert.Nil(t, capt)
	SetCustomStore(preStore)
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
	verifyValue, err := globalStore.Get(idkey, false)
	assert.Nil(t, err)
	fmt.Printf("idKey: %s, verifyValue: %s", idkey, verifyValue)
	if VerifyCaptcha(idkey, verifyValue) {
		t.Log(idkey, verifyValue)
	} else {
		t.Error("verify captcha content is failed.")
	}

	VerifyCaptcha("", "")
	VerifyCaptcha("dsafasf", "ddd")
}

func TestVerifyCaptchaOnFailed(t *testing.T) {

	preStore := globalStore
	s := store.NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6378",
	}), "test:", 5*time.Minute)
	SetCustomStore(s)
	assert.False(t, VerifyCaptcha("sss", "32321"))
	SetCustomStore(preStore)
}

func TestVerifyCaptchaV2(t *testing.T) {
	idkey, _, err := GenerateCaptchaV2("", configD)
	assert.Nil(t, err)
	verifyValue, err := globalStore.Get(idkey, false)
	assert.Nil(t, err)
	fmt.Printf("idKey: %s, verifyValue: %s", idkey, verifyValue)
	result, err := VerifyCaptchaV2(idkey, verifyValue)
	assert.Nil(t, err)
	if result {
		t.Log(idkey, verifyValue)
	} else {
		t.Error("verify captcha content is failed.")
	}

	VerifyCaptchaV2("", "")
	VerifyCaptchaV2("dsafasf", "ddd")
}

func TestVerifyCaptchaV2OnFailed(t *testing.T) {

	preStore := globalStore
	s := store.NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6378",
	}), "test:", 5*time.Minute)
	SetCustomStore(s)
	result, err := VerifyCaptchaV2("sss", "32321")
	assert.NotNil(t, err)
	assert.False(t, result)
	SetCustomStore(preStore)
}

func TestVerifyCaptchaAndIsClear(t *testing.T) {

	idkey, _ := GenerateCaptcha("", configD)
	verifyValue, err := globalStore.Get(idkey, false)
	assert.Nil(t, err)
	fmt.Printf("idKey: %s, verifyValue: %s", idkey, verifyValue)
	if VerifyCaptchaAndIsClear(idkey, verifyValue, true) {
		t.Log(idkey, verifyValue)
	} else {
		t.Error("verify captcha content is failed.")
	}

	VerifyCaptcha("", "")
	VerifyCaptcha("dsafasf", "ddd")
}

func TestVerifyCaptchaAndIsClearOnFailed(t *testing.T) {

	preStore := globalStore
	s := store.NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6378",
	}), "test:", 5*time.Minute)
	SetCustomStore(s)
	assert.False(t, VerifyCaptchaAndIsClear("sss", "wded", true))
	SetCustomStore(preStore)
}

func TestVerifyCaptchaAndIsClearV2(t *testing.T) {

	idkey, _ := GenerateCaptcha("", configD)
	verifyValue, err := globalStore.Get(idkey, false)
	assert.Nil(t, err)
	fmt.Printf("idKey: %s, verifyValue: %s", idkey, verifyValue)
	result, err := VerifyCaptchaAndIsClearV2(idkey, verifyValue, true)
	assert.Nil(t, err)
	if result {
		t.Log(idkey, verifyValue)
	} else {
		t.Error("verify captcha content is failed.")
	}

	result, err = VerifyCaptchaV2("", "")
	assert.Nil(t, err)
	assert.False(t, result)
	result, err = VerifyCaptchaV2("dsafasf", "ddd")
	assert.Nil(t, err)
	assert.False(t, result)
}

func TestVerifyCaptchaAndIsClearV2OnFailed(t *testing.T) {

	preStore := globalStore
	s := store.NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6378",
	}), "test:", 5*time.Minute)
	SetCustomStore(s)
	result, err := VerifyCaptchaAndIsClearV2("sss", "dddd", true)
	assert.NotNil(t, err)
	assert.False(t, result)
	SetCustomStore(preStore)
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
	// has no permission must failed
	if runtime.GOOS == "windows" {
		assert.Nil(t, err)
	} else {
		assert.NotNil(t, err)
	}
}

func TestSetCustomStore(t *testing.T) {
	s := store.NewMemoryStore(1000, 10*time.Minute)
	SetCustomStore(s)
	assert.Equal(t, globalStore, s)
}
