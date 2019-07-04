package base64Captcha

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"
)

//randText create random text. 生成随机文本.
func randText(num int, sourceChars string) string {
	textNum := len(sourceChars)
	text := ""
	r := randSeed()
	for i := 0; i < num; i++ {
		text = text + string(sourceChars[r.Intn(textNum)])
	}
	return text
}

//RandArithmetic create random arithmetic equation and result.
//穿件计算公式和返回结果
func randArithmetic() (question, result string) {
	operators := []string{"+", "-", "x"}
	var mathResult int32
	r := randSeed()
	switch operators[r.Int31n(3)] {
	case "+":
		a := r.Int31n(100)
		b := r.Int31n(100)
		question = fmt.Sprintf("%d+%d=?", a, b)
		mathResult = a + b
	case "x":
		a := r.Int31n(10)
		b := r.Int31n(10)
		question = fmt.Sprintf("%dx%d=?", a, b)
		mathResult = a * b
	default:
		a := r.Int31n(100)
		b := r.Int31n(100)
		if a > b {
			question = fmt.Sprintf("%d-%d=?", a, b)
			mathResult = a - b
		} else {
			question = fmt.Sprintf("%d-%d=?", b, a)
			mathResult = b - a
		}
	}
	result = fmt.Sprintf("%d", mathResult)
	return
}

//Random get random in min between max. 生成指定大小的随机数.
func random(min int64, max int64) float64 {

	r := randSeed()
	if max <= min {
		panic(fmt.Sprintf("invalid range %d >= %d", max, min))
	}
	decimal := r.Float64()

	if max <= 0 {
		return (float64(r.Int63n((min*-1)-(max*-1))+(max*-1)) + decimal) * -1
	}
	if min < 0 && max > 0 {
		if r.Int()%2 == 0 {
			return float64(r.Int63n(max)) + decimal
		}
		return (float64(r.Int63n(min*-1)) + decimal) * -1
	}
	return float64(r.Int63n(max-min)+min) + decimal
}

//randDeepColor get random deep color. 随机生成深色系.
func randDeepColor() color.RGBA {
	r := randSeed()

	randColor := randColor()

	increase := float64(30 + r.Intn(255))

	red := math.Abs(math.Min(float64(randColor.R)-increase, 255))

	green := math.Abs(math.Min(float64(randColor.G)-increase, 255))
	blue := math.Abs(math.Min(float64(randColor.B)-increase, 255))

	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//randLightColor get random ligth color. 随机生成浅色.
func randLightColor() color.RGBA {
	r := randSeed()

	red := r.Intn(55) + 200
	green := r.Intn(55) + 200
	blue := r.Intn(55) + 200

	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//randColor get random color. 生成随机颜色.
func randColor() color.RGBA {
	r := randSeed()

	red := r.Intn(255)
	green := r.Intn(255)
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

func randSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
