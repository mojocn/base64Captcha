package base64Captcha

import (
	"image/color"
	"strings"

	"github.com/golang/freetype/truetype"
)

// DriverString captcha config for captcha-engine-characters.
type DriverString struct {
	// Height png height in pixel.
	Height int

	// Width Captcha png width in pixel.
	Width int

	//NoiseCount text noise count.
	NoiseCount int

	//ShowLineOptions := OptionShowHollowLine | OptionShowSlimeLine | OptionShowSineLine .
	ShowLineOptions int

	//Length random string length.
	Length int

	//Source is a unicode which is the rand string from.
	Source string

	//BgColor captcha image background color (optional)
	BgColor *color.RGBA

	//fontsStorage font storage (optional)
	fontsStorage FontsStorage

	//Fonts loads by name see fonts.go's comment
	Fonts      []string
	fontsArray []*truetype.Font

	//MinFontSize minimum font size for text clarity (optional, default: calculated based on height)
	MinFontSize int

	//MaxFontSize maximum font size for text variation (optional, default: calculated based on height)  
	MaxFontSize int

	//Bold renders text in bold for better clarity (optional, default: false)
	Bold bool
}

// NewDriverString creates driver
func NewDriverString(height int, width int, noiseCount int, showLineOptions int, length int, source string, bgColor *color.RGBA, fontsStorage FontsStorage, fonts []string) *DriverString {
	if fontsStorage == nil {
		fontsStorage = DefaultEmbeddedFonts
	}

	tfs := []*truetype.Font{}
	for _, fff := range fonts {
		tf := fontsStorage.LoadFontByName("fonts/" + fff)
		tfs = append(tfs, tf)
	}

	if len(tfs) == 0 {
		tfs = fontsAll
	}

	// Calculate reasonable font size defaults based on image height
	minFontSize := height * 3 / 5  // 60% of height as minimum
	maxFontSize := height * 4 / 5  // 80% of height as maximum
	
	// Ensure minimum readability for small captchas
	if minFontSize < 16 {
		minFontSize = 16
	}
	if maxFontSize < minFontSize + 4 {
		maxFontSize = minFontSize + 4
	}

	return &DriverString{Height: height, Width: width, NoiseCount: noiseCount, ShowLineOptions: showLineOptions, Length: length, Source: source, BgColor: bgColor, fontsStorage: fontsStorage, fontsArray: tfs, Fonts: fonts, MinFontSize: minFontSize, MaxFontSize: maxFontSize}
}

// ConvertFonts loads fonts by names
func (d *DriverString) ConvertFonts() *DriverString {
	if d.fontsStorage == nil {
		d.fontsStorage = DefaultEmbeddedFonts
	}

	tfs := []*truetype.Font{}
	for _, fff := range d.Fonts {
		tf := d.fontsStorage.LoadFontByName("fonts/" + fff)
		tfs = append(tfs, tf)
	}
	if len(tfs) == 0 {
		tfs = fontsAll
	}

	d.fontsArray = tfs

	// Initialize font sizes if not set
	if d.MinFontSize == 0 || d.MaxFontSize == 0 {
		d.MinFontSize = d.Height * 3 / 5  // 60% of height as minimum
		d.MaxFontSize = d.Height * 4 / 5  // 80% of height as maximum
		
		// Ensure minimum readability for small captchas
		if d.MinFontSize < 16 {
			d.MinFontSize = 16
		}
		if d.MaxFontSize < d.MinFontSize + 4 {
			d.MaxFontSize = d.MinFontSize + 4
		}
	}

	return d
}

// GenerateIdQuestionAnswer creates id,content and answer
func (d *DriverString) GenerateIdQuestionAnswer() (id, content, answer string) {
	id = RandomId()
	content = RandText(d.Length, d.Source)
	return id, content, content
}

// DrawCaptcha draws captcha item
func (d *DriverString) DrawCaptcha(content string) (item Item, err error) {

	var bgc color.RGBA
	if d.BgColor != nil {
		bgc = *d.BgColor
	} else {
		bgc = RandLightColor()
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
		source := TxtNumbers + TxtAlphabet + ",.[]<>"
		noise := RandText(d.NoiseCount, strings.Repeat(source, d.NoiseCount))
		err = itemChar.drawNoise(noise, d.fontsArray)
		if err != nil {
			return
		}
	}

	//draw content
	err = itemChar.drawTextWithFontSize(content, d.fontsArray, d.MinFontSize, d.MaxFontSize, d.Bold)
	if err != nil {
		return
	}

	return itemChar, nil
}
