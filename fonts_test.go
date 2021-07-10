package base64Captcha

import (
	"testing"
)

// sources:
// fonts/3Dumb.ttf (142.224kB)
// fonts/ApothecaryFont.ttf (62.08kB)
// fonts/Comismsh.ttf (80.132kB)
// fonts/DENNEthree-dee.ttf (83.188kB)
// fonts/DeborahFancyDress.ttf (32.52kB)
// fonts/Flim-Flam.ttf (140.576kB)
// fonts/RitaSmith.ttf (31.24kB)
// fonts/actionj.ttf (34.944kB)
// fonts/chromohv.ttf (45.9kB)
// fonts/readme.md (162B)
// fonts/wqy-microhei.ttc (5.177MB)

func Test_loadFontByName(t *testing.T) {
	f := DefaultEmbeddedFonts.LoadFontByName("fonts/wqy-microhei.ttc")
	if f == nil {
		t.Error("failed")
	}

	defer recoverPanic(t)
	f = DefaultEmbeddedFonts.LoadFontByName("fonts/readme.md")

}
func recoverPanic(t *testing.T) {
	r := recover()
	if r == nil {
		t.Error("not trigger panic")
	}
}

func Test_loadFontsByNames(t *testing.T) {

	fs := DefaultEmbeddedFonts.LoadFontsByNames([]string{"fonts/chromohv.ttf", "fonts/RitaSmith.ttf"})
	if len(fs) != 2 {
		t.Error("failed")
	}
	defer recoverPanic(t)
	DefaultEmbeddedFonts.LoadFontsByNames([]string{"fonts/actionj.txxxxxtf"})
}

func Test_randFontFrom(t *testing.T) {
	f := randFontFrom(fontsAll)
	if f == nil {
		t.Error("failed")
	}
}
