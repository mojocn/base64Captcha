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
	"zh":     []int{0x4e00, 0x9fa5},
	"ko":     []int{12593, 12686},
	"jp":     []int{12449, 12531}, //[]int{12353, 12435}
	"ru":     []int{1025, 1169},
	"th":     []int{0x0e00, 0x0e7f},
	"greek":  []int{0x0380, 0x03ff},
	"arabic": []int{0x0600, 0x06ff},
	"hebrew": []int{0x0590, 0x05ff},
	//"emotion": []int{0x1f601, 0x1f64f},
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

//DriverLanguage generates language unicode by lanuage
type DriverLanguage struct {
	// Height png height in pixel.
	// 图像验证码的高度像素.
	Height int
	// Width Captcha png width in pixel.
	// 图像验证码的宽度像素
	Width int

	//NoiseCount text noise count.
	NoiseCount int

	ShowLineOptions int
	//CaptchaRunePairs make a list of rune for Captcha random selection.
	// 随机字符串可选内容

	// Length Default number of digits in captcha solution.
	// 默认数字验证长度6.
	Length int

	//BgColor captcha image background color (optional)
	//背景颜色
	BgColor      *color.RGBA
	Fonts        []*truetype.Font
	LanguageCode string
}

//NewDriverLanguage creates a driver
func NewDriverLanguage(height int, width int, noiseCount int, showLineOptions int, length int, bgColor *color.RGBA, fonts []*truetype.Font, languageCode string) *DriverLanguage {
	return &DriverLanguage{Height: height, Width: width, NoiseCount: noiseCount, ShowLineOptions: showLineOptions, Length: length, BgColor: bgColor, Fonts: fonts, LanguageCode: languageCode}
}

//GenerateQuestionAnswer creates content and answer
func (d *DriverLanguage) GenerateQuestionAnswer() (content, answer string) {
	content = generateRandomRune(d.Length, d.LanguageCode)
	return content, content
}

//GenerateItem creates item
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
	err = itemChar.drawText(content, []*truetype.Font{fontChinese})
	if err != nil {
		return
	}

	return itemChar, nil
}
