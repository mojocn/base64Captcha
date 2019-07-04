package base64Captcha

import (
	"testing"
)

func TestRandomId(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(randomDigits(5))
	}
}
func TestRandomDigits(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(randomDigits(5))
	}
}
func TestParseDigitsToString(t *testing.T) {
	for i := 0; i < 10; i++ {
		byss := randomDigits(5)
		t.Log(byss)
		bsssstring := parseDigitsToString(byss)
		t.Log(bsssstring)
	}
}
