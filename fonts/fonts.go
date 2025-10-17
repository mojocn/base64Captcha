package fonts

import (
	"github.com/mojocn/base64Captcha"
)

func init() {
	base64Captcha.DefaultEmbeddedFonts = base64Captcha.NewEmbeddedFontsStorage(defaultEmbeddedFontsFS)
	var fontsSimple = base64Captcha.DefaultEmbeddedFonts.LoadFontsByNames([]string{
		"3Dumb",
		"ApothecaryFont",
		"Comismsh",
		"DENNEthree-dee",
		"DeborahFancyDress",
		"Flim-Flam",
		"RitaSmith",
		"actionj",
		"chromohv",
	})
	base64Captcha.FontChinese = base64Captcha.DefaultEmbeddedFonts.LoadFontByName("wqy-microhei")
	base64Captcha.FontsAll = append(fontsSimple, base64Captcha.FontChinese)
}
