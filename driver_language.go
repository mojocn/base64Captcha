package base64Captcha

import (
	"github.com/golang/freetype/truetype"
	"image/color"
	"math/rand"
)

//https://en.wikipedia.org/wiki/Unicode_block
var langMap = map[string][]int{
	//"zh-CN": []int{19968, 40869},
	"latin":  []int{0x0000, 0x007f},
	"zh-CN":  []int{0x4e00, 0x9fa5},
	"ko":     []int{12593, 12686},
	"jp":     []int{12449, 12531}, //[]int{12353, 12435}
	"ru":     []int{1025, 1169},
	"th":     []int{0x0e00, 0x0e7f},
	"greek":  []int{0x0380, 0x03ff},
	"arabic": []int{0x0600, 0x06ff},
	"Hebrew": []int{0x0590, 0x05ff},
}

func generateRandomRune(size int, code string) string {
	lang, ok := langMap[code]
	if !ok {
		panic("can not find code")
	}
	start := lang[0]
	end := lang[1]
	randRune := make([]rune, size)
	for i := range randRune {
		idx := rand.Intn((end - start)) + start
		randRune[i] = rune(idx)
	}
	return string(randRune)
}

func NewDriverLanguage(driverString DriverString, languageCode string) *DriverLanguage {
	return &DriverLanguage{DriverString: driverString, LanguageCode: languageCode}
}

type DriverLanguage struct {
	DriverString
	LanguageCode string
}

func (d *DriverLanguage) GenerateQuestionAnswer() (content, answer string) {
	content = generateRandomRune(d.Length, d.LanguageCode)
	return content, content
}

func (d *DriverLanguage) GenerateItem(content string) (item Item, err error) {
	var bgc color.RGBA
	if d.BgColor != nil {
		bgc = *d.BgColor
	} else {
		bgc = randLightColor()
	}
	itemChar := NewItemChar(d.Width, d.Height, bgc)

	//draw hollow line
	if d.ShowLineOptions&OptionShowHollowLine == OptionShowHollowLine {
		itemChar.drawHollowLine()
	}

	//draw slime line
	if d.ShowLineOptions&OptionShowSlimeLine == OptionShowSlimeLine {
		itemChar.drawSlimLine(3)
	}

	//draw sine line
	if d.ShowLineOptions&OptionShowSineLine == OptionShowSineLine {
		itemChar.drawSineLine()
	}

	//draw noise
	if d.NoiseCount > 0 {
		noise := randText(d.NoiseCount, TxtNumbers+TxtAlphabet+",.[]<>")
		err = itemChar.drawNoise(noise, fontsAll)
		if err != nil {
			return
		}
	}

	//draw content
	//use font that match your language
	err = itemChar.DrawText(content, []*truetype.Font{fontChinese})
	if err != nil {
		return
	}

	return itemChar, nil
}
