// load fonts without chinese font
package defaults

import (
	"github.com/mojocn/base64Captcha"
	dennnethree_dee "github.com/mojocn/base64Captcha/fonts/DENNEthree-dee"
	"github.com/mojocn/base64Captcha/fonts/actionj"
	"github.com/mojocn/base64Captcha/fonts/apothecary"
	"github.com/mojocn/base64Captcha/fonts/chromohv"
	"github.com/mojocn/base64Captcha/fonts/comismsh"
	deborah_fancydress "github.com/mojocn/base64Captcha/fonts/deborahFancydress"
	flim_flam "github.com/mojocn/base64Captcha/fonts/flim-flam"
	rita_smith "github.com/mojocn/base64Captcha/fonts/rita-smith"
	three_dumb "github.com/mojocn/base64Captcha/fonts/three-dumb"
)

// defaultEmbeddedFontsFS Built-in font storage.
var defaultEmbeddedFontsFS = map[string][]byte{
	"3Dumb":             three_dumb.FontBytes,
	"ApothecaryFont":    apothecary.FontBytes,
	"Comismsh":          comismsh.FontBytes,
	"DENNEthree-dee":    dennnethree_dee.FontBytes,
	"DeborahFancyDress": deborah_fancydress.FontBytes,
	"Flim-Flam":         flim_flam.FontBytes,
	"RitaSmith":         rita_smith.FontBytes,
	"actionj":           actionj.FontBytes,
	"chromohv":          chromohv.FontBytes,
}

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
	base64Captcha.FontsAll = append(fontsSimple, base64Captcha.FontChinese)
}
