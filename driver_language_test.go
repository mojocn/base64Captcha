package base64Captcha

import (
	"github.com/golang/freetype/truetype"
	"testing"
)

func TestDriverLanguage_DrawCaptcha(t *testing.T) {
	ds := NewDriverLanguage(80, 240, 5, OptionShowSineLine|OptionShowSlimeLine|OptionShowHollowLine, 5, nil, []*truetype.Font{fontChinese}, "emotion")

	for i := 0; i < 40; i++ {
		_, q, _ := ds.GenerateIdQuestionAnswer()
		item, err := ds.DrawCaptcha(q)
		if err != nil {
			t.Error(err)
		}
		itemWriteFile(item, "_builds", randomId(), "png")
	}
}
