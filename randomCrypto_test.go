package base64Captcha

import (
	"testing"
)

func TestRandomId(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomDigits(5))
	}
}
func TestRandomDigits(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomDigits(5))
	}
}
func TestParseDigitsToString(t *testing.T) {
	for i := 0; i < 10; i++ {
		byss := RandomDigits(5)
		t.Log(byss)
		bsssstring := ParseDigitsToString(byss)
		t.Log(bsssstring)
	}
}
