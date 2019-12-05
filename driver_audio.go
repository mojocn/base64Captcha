// Copyright 2017 Eric Zhou. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base64Captcha

//driverAudio captcha config for captcha-engine-audio.
type driverAudio struct {
	// Length Default number of digits in captcha solution.
	Length int
	// Language possible values for lang are "en", "ja", "ru", "zh".
	Language string
}

func NewDriverAudio(length int, language string) *driverAudio {
	return &driverAudio{Length: length, Language: language}
}

func (d *driverAudio) GenerateItem(content string) (item Item, err error) {
	digits := stringToFakeByte(content)
	audio := newAudio("", digits, d.Language)
	return audio, nil
}
func (d *driverAudio) GenerateQuestionAnswer() (q, a string) {
	digits := randomDigits(d.Length)
	a = parseDigitsToString(digits)
	return a, a
}
