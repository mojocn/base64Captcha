package base64Captcha

import "testing"

func TestAsset(t *testing.T) {

	for idx, value := range _bindata {

		t.Log(value, idx)

		ass, _ := value()

		t.Log(ass.info.Name(), ass.info.Size(), ass.info.Mode(), ass.info.ModTime(), ass.info.Sys(), ass.info.IsDir())

	}

}
