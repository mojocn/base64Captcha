package base64Captcha

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math"
	"math/rand"
)

var trueTypeFontFamilys = readFontsToSliceOfTrueTypeFonts()

//CaptchaImageChar captcha-engine-char return type.
type CaptchaImageChar struct {
	CaptchaItem
	nrgba   *image.NRGBA
	Complex int
}

//ConfigCharacter captcha config for captcha-engine-characters.
type ConfigCharacter struct {
	// Height png height in pixel.
	// 图像验证码的高度像素.
	Height int
	// Width Captcha png width in pixel.
	// 图像验证码的宽度像素
	Width int
	//Mode : base64captcha.CaptchaModeNumber=0, base64captcha.CaptchaModeAlphabet=1, base64captcha.CaptchaModeArithmetic=2, base64captcha.CaptchaModeNumberAlphabet=3.
	Mode int
	//IsUseSimpleFont is use simply font(...base64Captcha/fonts/RitaSmith.ttf).
	IsUseSimpleFont bool
	//ComplexOfNoiseText text noise count.
	ComplexOfNoiseText int
	//ComplexOfNoiseDot dot noise count.
	ComplexOfNoiseDot int
	//IsShowHollowLine is show hollow line.
	IsShowHollowLine bool
	//IsShowNoiseDot is show noise dot.
	IsShowNoiseDot bool
	//IsShowNoiseText is show noise text.
	IsShowNoiseText bool
	//IsShowSlimeLine is show slime line.
	IsShowSlimeLine bool
	//IsShowSineLine is show sine line.
	IsShowSineLine bool

	// CaptchaLen Default number of digits in captcha solution.
	// 默认数字验证长度6.
	CaptchaLen int
}
type point struct {
	X int
	Y int
}

//newCaptchaImage new blank captchaImage context.
//新建一个图片对象.
func newCaptchaImage(width int, height int, bgColor color.RGBA) (cImage *CaptchaImageChar, err error) {
	m := image.NewNRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)
	cImage = &CaptchaImageChar{}
	cImage.nrgba = m
	cImage.ImageHeight = height
	cImage.ImageWidth = width
	err = nil
	return
}

//drawHollowLine draw strong and bold white line.
//添加一个较粗的空白直线
func (captcha *CaptchaImageChar) drawHollowLine() *CaptchaImageChar {

	first := captcha.ImageWidth / 20
	end := first * 19

	lineColor := color.RGBA{R: 245, G: 250, B: 251, A: 255}

	x1 := float64(rand.Intn(first))
	//y1 := float64(rand.Intn(y)+y);

	x2 := float64(rand.Intn(first) + end)

	multiple := float64(rand.Intn(5)+3) / float64(5)
	if int(multiple*10)%3 == 0 {
		multiple = multiple * -1.0
	}

	w := captcha.ImageHeight / 20

	for ; x1 < x2; x1++ {

		y := math.Sin(x1*math.Pi*multiple/float64(captcha.ImageWidth)) * float64(captcha.ImageHeight/3)

		if multiple < 0 {
			y = y + float64(captcha.ImageHeight/2)
		}
		captcha.nrgba.Set(int(x1), int(y), lineColor)

		for i := 0; i <= w; i++ {
			captcha.nrgba.Set(int(x1), int(y)+i, lineColor)
		}
	}

	return captcha
}

//drawSineLine draw a sine line.
//画一条正弦曲线.
func (captcha *CaptchaImageChar) drawSineLine() *CaptchaImageChar {
	var py float64

	//振幅
	a := rand.Intn(captcha.ImageHeight / 2)

	//Y轴方向偏移量
	b := random(int64(-captcha.ImageHeight/4), int64(captcha.ImageHeight/4))

	//X轴方向偏移量
	f := random(int64(-captcha.ImageHeight/4), int64(captcha.ImageHeight/4))
	// 周期
	var t float64
	if captcha.ImageHeight > captcha.ImageWidth/2 {
		t = random(int64(captcha.ImageWidth/2), int64(captcha.ImageHeight))
	} else if captcha.ImageHeight == captcha.ImageWidth/2 {
		t = float64(captcha.ImageHeight)
	} else {
		t = random(int64(captcha.ImageHeight), int64(captcha.ImageWidth/2))
	}
	w := float64((2 * math.Pi) / t)

	// 曲线横坐标起始位置
	px1 := 0
	px2 := int(random(int64(float64(captcha.ImageWidth)*0.8), int64(captcha.ImageWidth)))

	c := color.RGBA{R: uint8(rand.Intn(150)), G: uint8(rand.Intn(150)), B: uint8(rand.Intn(150)), A: uint8(255)}

	for px := px1; px < px2; px++ {
		if w != 0 {
			py = float64(a)*math.Sin(w*float64(px)+f) + b + (float64(captcha.ImageWidth) / float64(5))
			i := captcha.ImageHeight / 5
			for i > 0 {
				captcha.nrgba.Set(px+i, int(py), c)
				//fmt.Println(px + i,int(py) )
				i--
			}
		}
	}

	return captcha
}

//drawSlimLine draw n slim-random-color lines.
//画n条随机颜色的细线
func (captcha *CaptchaImageChar) drawSlimLine(num int) *CaptchaImageChar {

	first := captcha.ImageWidth / 10
	end := first * 9

	y := captcha.ImageHeight / 3

	for i := 0; i < num; i++ {

		point1 := point{X: rand.Intn(first), Y: rand.Intn(y)}
		point2 := point{X: rand.Intn(first) + end, Y: rand.Intn(y)}

		if i%2 == 0 {
			point1.Y = rand.Intn(y) + y*2
			point2.Y = rand.Intn(y)
		} else {
			point1.Y = rand.Intn(y) + y*(i%2)
			point2.Y = rand.Intn(y) + y*2
		}

		captcha.drawBeeline(point1, point2, randDeepColor())

	}
	return captcha
}

func (captcha *CaptchaImageChar) drawBeeline(point1 point, point2 point, lineColor color.RGBA) {
	dx := math.Abs(float64(point1.X - point2.X))

	dy := math.Abs(float64(point2.Y - point1.Y))
	sx, sy := 1, 1
	if point1.X >= point2.X {
		sx = -1
	}
	if point1.Y >= point2.Y {
		sy = -1
	}
	err := dx - dy
	for {
		captcha.nrgba.Set(point1.X, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X+1, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X-1, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X+2, point1.Y, lineColor)
		captcha.nrgba.Set(point1.X-2, point1.Y, lineColor)
		if point1.X == point2.X && point1.Y == point2.Y {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			point1.X += sx
		}
		if e2 < dx {
			err += dx
			point1.Y += sy
		}
	}
}

//drawNoise draw noise dots.
//画干扰点.
func (captcha *CaptchaImageChar) drawNoise(complex int) *CaptchaImageChar {
	density := 18
	if complex == CaptchaComplexLower {
		density = 28
	} else if complex == CaptchaComplexMedium {
		density = 18
	} else if complex == CaptchaComplexHigh {
		density = 8
	}
	maxSize := (captcha.ImageHeight * captcha.ImageWidth) / density

	for i := 0; i < maxSize; i++ {

		rw := rand.Intn(captcha.ImageWidth)
		rh := rand.Intn(captcha.ImageHeight)

		captcha.nrgba.Set(rw, rh, randColor())
		size := rand.Intn(maxSize)
		if size%3 == 0 {
			captcha.nrgba.Set(rw+1, rh+1, randColor())
		}
	}
	return captcha
}

//drawTextNoise draw noises which are single character.
//画文字噪点.
func (captcha *CaptchaImageChar) drawTextNoise(complex int, isSimpleFont bool) error {
	density := 1500
	if complex == CaptchaComplexLower {
		density = 2000
	} else if complex == CaptchaComplexMedium {
		density = 1500
	} else if complex == CaptchaComplexHigh {
		density = 1000
	}

	maxSize := (captcha.ImageHeight * captcha.ImageWidth) / density

	//r := rand.New(rand.NewSource(time.Now().UnixNano()))

	c := freetype.NewContext()
	c.SetDPI(imageStringDpi)

	c.SetClip(captcha.nrgba.Bounds())
	c.SetDst(captcha.nrgba)
	c.SetHinting(font.HintingFull)
	rawFontSize := float64(captcha.ImageHeight) / (1 + float64(rand.Intn(7))/float64(10))

	for i := 0; i < maxSize; i++ {

		rw := rand.Intn(captcha.ImageWidth)
		rh := rand.Intn(captcha.ImageHeight)

		text := randText(1, TxtNumbers+TxtAlphabet)
		fontSize := rawFontSize/2 + float64(rand.Intn(5))

		c.SetSrc(image.NewUniform(randLightColor()))
		c.SetFontSize(fontSize)

		if isSimpleFont {
			c.SetFont(trueTypeFontFamilys[0])
		} else {
			f := randFontFamily()
			c.SetFont(f)
		}

		pt := freetype.Pt(rw, rh)

		if _, err := c.DrawString(text, pt); err != nil {
			log.Println(err)
		}
	}
	return nil
}

//drawText draw captcha string to image.把文字写入图像验证码
func (captcha *CaptchaImageChar) drawText(text string, isSimpleFont bool) error {
	c := freetype.NewContext()
	c.SetDPI(imageStringDpi)

	c.SetClip(captcha.nrgba.Bounds())
	c.SetDst(captcha.nrgba)
	c.SetHinting(font.HintingFull)

	fontWidth := captcha.ImageWidth / len(text)

	for i, s := range text {

		fontSize := float64(captcha.ImageHeight) / (1 + float64(rand.Intn(7))/float64(9))

		c.SetSrc(image.NewUniform(randDeepColor()))
		c.SetFontSize(fontSize)

		if isSimpleFont {
			c.SetFont(trueTypeFontFamilys[0])
		} else {
			f := randFontFamily()
			c.SetFont(f)
		}

		x := int(fontWidth)*i + int(fontWidth)/int(fontSize)

		y := 5 + rand.Intn(captcha.ImageHeight/2) + int(fontSize/2)

		pt := freetype.Pt(x, y)

		if _, err := c.DrawString(string(s), pt); err != nil {
			log.Println(err)
		}
		//pt.Y += c.pointToFixed(*size * *spacing)
		//pt.X += c.pointToFixed(*size);
	}
	return nil

}

//EngineCharCreate create captcha with config struct.
func EngineCharCreate(config ConfigCharacter) *CaptchaImageChar {

	captchaImage, err := newCaptchaImage(config.Width, config.Height, randLightColor())

	//背景有像素点干扰
	if config.IsShowNoiseDot {
		captchaImage.drawNoise(config.ComplexOfNoiseDot)
	}

	//波浪线       比较丑
	if config.IsShowHollowLine {
		captchaImage.drawHollowLine()
	}
	//背景有文字干扰
	if config.IsShowNoiseText {
		captchaImage.drawTextNoise(config.ComplexOfNoiseText, config.IsUseSimpleFont)
	}

	//画 细直线 (n 条)
	if config.IsShowSlimeLine {
		captchaImage.drawSlimLine(3)
	}

	//画 多个小波浪线
	if config.IsShowSineLine {
		captchaImage.drawSineLine()
	}
	var captchaContent string

	switch config.Mode {
	case CaptchaModeAlphabet:
		captchaContent = randText(config.CaptchaLen, TxtAlphabet)
		captchaImage.VerifyValue = captchaContent
	case CaptchaModeArithmetic:
		captchaContent, captchaImage.VerifyValue = randArithmetic()

	case CaptchaModeNumber:
		captchaContent = randText(config.CaptchaLen, TxtNumbers)
		captchaImage.VerifyValue = captchaContent
	default:
		captchaContent = randText(config.CaptchaLen, TxtSimpleCharaters)
		captchaImage.VerifyValue = captchaContent
	}
	//写入string
	captchaImage.drawText(captchaContent, config.IsUseSimpleFont)
	captchaImage.Content = captchaContent
	//captchaImage.drawText(randText(4))

	if err != nil {
		fmt.Println(err)
	}

	return captchaImage
}

//BinaryEncodeing save captcha image to binary.
//保存图片到io.
func (captcha *CaptchaImageChar) BinaryEncodeing() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, captcha.nrgba); err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}

// WriteTo writes captcha image in PNG format into the given writer.
func (captcha *CaptchaImageChar) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(captcha.BinaryEncodeing())
	return int64(n), err
}
