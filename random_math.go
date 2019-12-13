package base64Captcha

import (
	"image/color"
	"math"
	"math/rand"
	"strings"
)

//RandText create random text.
func RandText(size int, sourceChars string) string {
	if size >= len(sourceChars) {
		sourceChars = strings.Repeat(sourceChars, size)
	}
	sourceRunes := []rune(sourceChars)
	sourceLength := len(sourceRunes)

	text := make([]rune, size)
	for i := range text {
		text[i] = sourceRunes[rand.Intn(sourceLength)]
	}
	return string(text)
}

//Random get random in min between max. 生成指定大小的随机数.
func random(min int64, max int64) float64 {
	return rand.Float64()*float64(max) - float64(min)
}

//RandDeepColor get random deep color. 随机生成深色系.
func RandDeepColor() color.RGBA {

	randColor := RandColor()

	increase := float64(30 + rand.Intn(255))

	red := math.Abs(math.Min(float64(randColor.R)-increase, 255))

	green := math.Abs(math.Min(float64(randColor.G)-increase, 255))
	blue := math.Abs(math.Min(float64(randColor.B)-increase, 255))

	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//RandLightColor get random ligth color. 随机生成浅色.
func RandLightColor() color.RGBA {
	red := rand.Intn(55) + 200
	green := rand.Intn(55) + 200
	blue := rand.Intn(55) + 200
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//RandColor get random color. 生成随机颜色.
func RandColor() color.RGBA {

	red := rand.Intn(255)
	green := rand.Intn(255)
	var blue int
	if (red + green) > 400 {
		blue = 0
	} else {
		blue = 400 - green - red
	}
	if blue > 255 {
		blue = 255
	}
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}
