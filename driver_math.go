package base64Captcha

import (
	"fmt"
	"image/color"
	"math/rand"
)

//DriverChar captcha config for captcha-engine-characters.
type DriverMath DriverString

func (d *DriverMath) GenerateQuestionAnswer() (question, answer string) {
	operators := []string{"+", "-", "x"}
	var mathResult int32
	switch operators[rand.Int31n(3)] {
	case "+":
		a := rand.Int31n(100)
		b := rand.Int31n(100)
		question = fmt.Sprintf("%d+%d=?", a, b)
		mathResult = a + b
	case "x":
		a := rand.Int31n(10)
		b := rand.Int31n(10)
		question = fmt.Sprintf("%dx%d=?", a, b)
		mathResult = a * b
	default:
		a := rand.Int31n(100)
		b := rand.Int31n(100)
		if a > b {
			question = fmt.Sprintf("%d-%d=?", a, b)
			mathResult = a - b
		} else {
			question = fmt.Sprintf("%d-%d=?", b, a)
			mathResult = b - a
		}
	}
	answer = fmt.Sprintf("%d", mathResult)
	return
}
func (d *DriverMath) GenerateItem(question string) (item Item, err error) {

	var bgc color.RGBA
	if d.BgColor != nil {
		bgc = *d.BgColor
	} else {
		bgc = randLightColor()
	}
	itemChar := NewItemChar(d.Width, d.Height, bgc)

	//波浪线 比较丑
	if d.ShowLineOptions&OptionShowHollowLine == OptionShowHollowLine {
		itemChar.drawHollowLine()
	}

	//背景有文字干扰
	if d.NoiseCount > 0 {
		noise := randText(d.NoiseCount, TxtNumbers)
		err = itemChar.drawNoise(noise, fontsAll)
		if err != nil {
			return
		}
	}

	//画 细直线 (n 条)
	if d.ShowLineOptions&OptionShowSlimeLine == OptionShowSlimeLine {
		itemChar.drawSlimLine(3)
	}

	//画 多个小波浪线
	if d.ShowLineOptions&OptionShowSineLine == OptionShowSineLine {
		itemChar.drawSineLine()
	}

	//draw question
	err = itemChar.DrawText(question, fontsAll)
	if err != nil {
		return
	}
	return itemChar, nil
}
