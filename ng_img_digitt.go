// Copyright 2017 Eric Zhou. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base64Captcha

import (
	"bytes"

	"image"
	"image/color"
	"image/png"
	"io"
	"math"
)

//ConfigDigit config for captcha-engine-digit.
type ConfigDigit struct {

	// Height png height in pixel.
	// 图像验证码的高度像素.
	Height int
	// Width Captcha png width in pixel.
	// 图像验证码的宽度像素
	Width int
	// DefaultLen Default number of digits in captcha solution.
	// 默认数字验证长度6.
	CaptchaLen int
	// MaxSkew max absolute skew factor of a single digit.
	// 图像验证码的最大干扰洗漱.
	MaxSkew float64
	// DotCount Number of background circles.
	// 图像验证码干扰圆点的数量.
	DotCount int
}

// CaptchaImageDigit digits captcha Struct
type CaptchaImageDigit struct {
	CaptchaItem
	*image.Paletted
	dotSize int
	rng     siprng
}

//EngineDigitsCreate create captcha by engine-digits with configuration.
func EngineDigitsCreate(id string, config ConfigDigit) *CaptchaImageDigit {
	digits := randomDigits(config.CaptchaLen)
	// Initialize PRNG.
	m := new(CaptchaImageDigit)
	//parse digits to string
	m.VerifyValue = parseDigitsToString(digits)

	m.rng.Seed(deriveSeed(imageSeedPurpose, id, digits))

	m.Paletted = image.NewPaletted(image.Rect(0, 0, config.Width, config.Height), m.getRandomPalette())
	m.calculateSizes(config.Width, config.Height, len(digits))
	// Randomly position captcha inside the image.
	maxx := config.Width - (m.ImageWidth+m.dotSize)*len(digits) - m.dotSize
	maxy := config.Height - m.ImageHeight - m.dotSize*2
	var border int
	if config.Width > config.Height {
		border = config.Height / 5
	} else {
		border = config.Width / 5
	}
	x := m.rng.Int(border, maxx-border)
	y := m.rng.Int(border, maxy-border)
	// Draw digits.
	for _, n := range digits {
		m.drawDigit(digitFontData[n], x, y)
		x += m.ImageWidth + m.dotSize
	}
	// Draw strike-through line.
	m.strikeThrough()
	// Apply wave distortion.
	m.distort(m.rng.Float(5, 10), m.rng.Float(100, 200))
	// Fill image with random circles.
	m.fillWithCircles(DotCount, m.dotSize)
	return m
}

func (m *CaptchaImageDigit) getRandomPalette() color.Palette {
	p := make([]color.Color, DotCount+1)
	// Transparent color.
	p[0] = color.RGBA{0xFF, 0xFF, 0xFF, 0x00}
	// Primary color.
	prim := color.RGBA{
		uint8(m.rng.Intn(129)),
		uint8(m.rng.Intn(129)),
		uint8(m.rng.Intn(129)),
		0xFF,
	}
	p[1] = prim
	// Circle colors.
	for i := 2; i <= DotCount; i++ {
		p[i] = m.randomBrightness(prim, 255)
	}
	return p
}

func (m *CaptchaImageDigit) calculateSizes(width, height, ncount int) {
	// Goal: fit all digits inside the image.
	var border int
	if width > height {
		border = height / 4
	} else {
		border = width / 4
	}
	// Convert everything to floats for calculations.
	w := float64(width - border*2)
	h := float64(height - border*2)
	// fw takes into account 1-dot spacing between digits.
	fw := float64(digitFontWidth + 1)
	fh := float64(digitFontHeight)
	nc := float64(ncount)
	// Calculate the width of a single digit taking into account only the
	// width of the image.
	nw := w / nc
	// Calculate the height of a digit from this width.
	nh := nw * fh / fw
	// Digit too high?
	if nh > h {
		// Fit digits based on height.
		nh = h
		nw = fw / fh * nh
	}
	// Calculate dot size.
	m.dotSize = int(nh / fh)
	if m.dotSize < 1 {
		m.dotSize = 1
	}
	// Save everything, making the actual width smaller by 1 dot to account
	// for spacing between digits.
	m.ImageWidth = int(nw) - m.dotSize
	m.ImageHeight = int(nh)
}

func (m *CaptchaImageDigit) drawHorizLine(fromX, toX, y int, colorIdx uint8) {
	for x := fromX; x <= toX; x++ {
		m.SetColorIndex(x, y, colorIdx)
	}
}

func (m *CaptchaImageDigit) drawCircle(x, y, radius int, colorIdx uint8) {
	f := 1 - radius
	dfx := 1
	dfy := -2 * radius
	xo := 0
	yo := radius

	m.SetColorIndex(x, y+radius, colorIdx)
	m.SetColorIndex(x, y-radius, colorIdx)
	m.drawHorizLine(x-radius, x+radius, y, colorIdx)

	for xo < yo {
		if f >= 0 {
			yo--
			dfy += 2
			f += dfy
		}
		xo++
		dfx += 2
		f += dfx
		m.drawHorizLine(x-xo, x+xo, y+yo, colorIdx)
		m.drawHorizLine(x-xo, x+xo, y-yo, colorIdx)
		m.drawHorizLine(x-yo, x+yo, y+xo, colorIdx)
		m.drawHorizLine(x-yo, x+yo, y-xo, colorIdx)
	}
}

func (m *CaptchaImageDigit) fillWithCircles(n, maxradius int) {
	maxx := m.Bounds().Max.X
	maxy := m.Bounds().Max.Y
	for i := 0; i < n; i++ {
		colorIdx := uint8(m.rng.Int(1, DotCount-1))
		r := m.rng.Int(1, maxradius)
		m.drawCircle(m.rng.Int(r, maxx-r), m.rng.Int(r, maxy-r), r, colorIdx)
	}
}

func (m *CaptchaImageDigit) strikeThrough() {
	maxx := m.Bounds().Max.X
	maxy := m.Bounds().Max.Y
	y := m.rng.Int(maxy/3, maxy-maxy/3)
	amplitude := m.rng.Float(5, 20)
	period := m.rng.Float(80, 180)
	dx := 2.0 * math.Pi / period
	for x := 0; x < maxx; x++ {
		xo := amplitude * math.Cos(float64(y)*dx)
		yo := amplitude * math.Sin(float64(x)*dx)
		for yn := 0; yn < m.dotSize; yn++ {
			r := m.rng.Int(0, m.dotSize)
			m.drawCircle(x+int(xo), y+int(yo)+(yn*m.dotSize), r/2, 1)
		}
	}
}

//写入数字 数字byte
func (m *CaptchaImageDigit) drawDigit(digit []byte, x, y int) {
	skf := m.rng.Float(-MaxSkew, MaxSkew)
	xs := float64(x)
	r := m.dotSize / 2
	y += m.rng.Int(-r, r)
	for yo := 0; yo < digitFontHeight; yo++ {
		for xo := 0; xo < digitFontWidth; xo++ {
			if digit[yo*digitFontWidth+xo] != digitFontBlackChar {
				continue
			}
			m.drawCircle(x+xo*m.dotSize, y+yo*m.dotSize, r, 1)
		}
		xs += skf
		x = int(xs)
	}
}

func (m *CaptchaImageDigit) distort(amplude float64, period float64) {
	w := m.Bounds().Max.X
	h := m.Bounds().Max.Y

	oldm := m.Paletted
	newm := image.NewPaletted(image.Rect(0, 0, w, h), oldm.Palette)

	dx := 2.0 * math.Pi / period
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			xo := amplude * math.Sin(float64(y)*dx)
			yo := amplude * math.Cos(float64(x)*dx)
			newm.SetColorIndex(x, y, oldm.ColorIndexAt(x+int(xo), y+int(yo)))
		}
	}
	m.Paletted = newm
}

func (m *CaptchaImageDigit) randomBrightness(c color.RGBA, max uint8) color.RGBA {
	minc := min3(c.R, c.G, c.B)
	maxc := max3(c.R, c.G, c.B)
	if maxc > max {
		return c
	}
	n := m.rng.Intn(int(max-maxc)) - int(minc)
	return color.RGBA{
		uint8(int(c.R) + n),
		uint8(int(c.G) + n),
		uint8(int(c.B) + n),
		uint8(c.A),
	}
}

func min3(x, y, z uint8) (m uint8) {
	m = x
	if y < m {
		m = y
	}
	if z < m {
		m = z
	}
	return
}

func max3(x, y, z uint8) (m uint8) {
	m = x
	if y > m {
		m = y
	}
	if z > m {
		m = z
	}
	return
}

// BinaryEncodeing encodes an image to PNG and returns
// the result as a byte slice.
func (m *CaptchaImageDigit) BinaryEncodeing() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, m.Paletted); err != nil {
		panic(err.Error())
	}
	return buf.Bytes()
}

// WriteTo writes captcha image in PNG format into the given writer.
func (m *CaptchaImageDigit) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(m.BinaryEncodeing())
	return int64(n), err
}
