package base64Captcha

import "testing"

func TestHandlerCaptchaGenerate(t *testing.T) {
	s := DefaultMemStore

	driver := &DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      10,
		ShowLineOptions: 10,
		Length:          10,
		Source:          "axclajsdlfkjalskjdglasdg",
		BgColor:         nil,
		Fonts:           nil,
	}

	c := NewCaptcha(driver, s)

	id, _, _, err := c.Generate()
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	t.Logf("id: %s", id)
}
