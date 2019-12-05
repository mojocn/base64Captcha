package base64Captcha

import (
	"github.com/golang/freetype/truetype"
	"image/color"
)

//DriverChar captcha config for captcha-engine-characters.
type DriverString struct {
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
	BgColor *color.RGBA
	Fonts   []*truetype.Font
}

func NewDriverString(height int, width int, noiseCount int, showLineOptions int, length int, bgColor *color.RGBA, fonts []*truetype.Font) *DriverString {
	return &DriverString{Height: height, Width: width, NoiseCount: noiseCount, ShowLineOptions: showLineOptions, Length: length, BgColor: bgColor, Fonts: fonts}
}

func (d *DriverString) GenerateQuestionAnswer() (content, answer string) {
	content = randText(d.Length, TxtAlphabet+TxtNumbers)
	return content, content
}
func (d *DriverString) GenerateItem(content string) (item Item, err error) {
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
	err = itemChar.DrawText(content, fontsAll)
	if err != nil {
		return
	}

	return itemChar, nil
}
