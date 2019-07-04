package base64Captcha

import "testing"

func TestAsset(t *testing.T) {

	for _, value := range _bindata {

		_, err := value()
		if err != nil {
			t.Error(err)
		}

	}

}
