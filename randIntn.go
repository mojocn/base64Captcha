package base64Captcha

import "math/rand"

func randIntn(n int) int {
	if n > 0 {
		return rand.Intn(n)
	}
	return 0
}
