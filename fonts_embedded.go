package base64Captcha

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type EmbeddedFontsStorage struct {
	data map[string][]byte
}

func (s *EmbeddedFontsStorage) LoadFontByName(name string) *truetype.Font {
	fontBytes, ok := s.data[name]
	if !ok {
		panic("font not found.")
	}

	//font file bytes to trueTypeFont
	trueTypeFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	return trueTypeFont
}

// LoadFontsByNames import fonts from dir.
// make the simple-font(RitaSmith.ttf) the first font of trueTypeFonts.
func (s *EmbeddedFontsStorage) LoadFontsByNames(assetFontNames []string) []*truetype.Font {
	fonts := make([]*truetype.Font, 0)
	for _, assetName := range assetFontNames {
		f := s.LoadFontByName(assetName)
		fonts = append(fonts, f)
	}
	return fonts
}

func NewEmbeddedFontsStorage(data map[string][]byte) *EmbeddedFontsStorage {
	return &EmbeddedFontsStorage{
		data: data,
	}
}
