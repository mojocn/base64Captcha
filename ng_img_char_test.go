package base64Captcha

import (
	"bytes"
	"image/color"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/freetype/truetype"
)

func TestEngineCharCreate(t *testing.T) {
	tc, _ := ioutil.TempDir("", "audio")
	defer os.Remove(tc)
	for i := 0; i < 16; i++ {
		configC.Mode = i % 5
		boooo := i%2 == 0
		configC.IsUseSimpleFont = boooo
		configC.IsShowSlimeLine = boooo
		configC.IsShowNoiseText = boooo
		configC.IsShowHollowLine = boooo
		configC.IsShowSineLine = boooo
		configC.IsShowNoiseDot = boooo

		if configC.Mode == CaptchaModeChinese {
			configC.UseCJKFonts = true
		} else {
			configC.UseCJKFonts = false
		}

		im := EngineCharCreate(configC)
		fileName := strings.Trim(im.Content, "/+-+=?")
		err := CaptchaWriteToFile(im, tc, fileName, "png")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestEngineCharUseSequencedCharacters(t *testing.T) {
	tc, _ := ioutil.TempDir("", "UseSequencedCharacters")
	defer os.Remove(tc)
	configSequencedCharacters := configC
	configSequencedCharacters.UseCJKFonts = true
	configSequencedCharacters.Mode = CaptchaModeUseSequencedCharacters
	configSequencedCharacters.SequencedCharacters = []string{
		"文件", "下载", "测试", "词组", "验证", "顺序",
	}
	configSequencedCharacters.CaptchaLen = 9
	im := EngineCharCreate(configSequencedCharacters)
	rep := im.Content
	for _, each := range configSequencedCharacters.SequencedCharacters {
		rep = strings.Replace(rep, each, "", -1)
	}
	if len([]rune(rep)) != 1 {
		t.Errorf("notgood: %v [rep=%v]", im.Content, rep)
	}
	fileName := strings.Trim(im.Content, "/+-+=?")
	err := CaptchaWriteToFile(im, tc, fileName, "png")
	if err != nil {
		t.Error(err)
	}
}

func TestEngineCharCreateShowNoiseText(t *testing.T) {
	tc, _ := ioutil.TempDir("", "TestEngineCharCreateShowNoiseText")
	defer os.Remove(tc)
	configC.Mode = CaptchaModeNumber
	configC.IsShowNoiseText = true
	random_int := 0x123345
	for _, complex := range []int{CaptchaComplexLower, CaptchaComplexMedium, CaptchaComplexHigh, random_int} {
		configC.ComplexOfNoiseText = complex
		im := EngineCharCreate(configC)
		fileName := strings.Trim(im.Content, "/+-+=?")
		err := CaptchaWriteToFile(im, tc, fileName, "png")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestEngineCharCreatePanics(t *testing.T) {
	tc, _ := ioutil.TempDir("", "TestEngineCharCreateShowNoiseText")
	defer os.Remove(tc)

	t.Run("TestCaptchaLen_EQ_0", func(t *testing.T) {
		configBad := configC
		configBad.CaptchaLen = 0
		defer func() {
			if err := recover(); err != nil {
				//good panic
				return
			}
		}()
		_ = EngineCharCreate(configBad)
		t.Error("shell not be here")
	})

	t.Run("TestCaptchaModeUseSequencedCharacters", func(t *testing.T) {
		configBad := configC
		configBad.CaptchaLen = 9
		configBad.Mode = CaptchaModeUseSequencedCharacters
		configBad.SequencedCharacters = nil
		defer func() {
			if err := recover(); err != nil {
				//good panic
				return
			}
		}()
		_ = EngineCharCreate(configBad)
		t.Error("shell not be here")
	})

}

func TestMath(t *testing.T) {
	for i := 0; i < 100; i++ {
		q, r := randArithmetic()
		t.Log(q, "--->", r)
	}
}

func Test_newCaptchaImage(t *testing.T) {
	type args struct {
		width   int
		height  int
		bgColor color.RGBA
	}
	tests := []struct {
		name       string
		args       args
		wantCImage *CaptchaImageChar
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCImage, err := newCaptchaImage(tt.args.width, tt.args.height, tt.args.bgColor)
			if (err != nil) != tt.wantErr {
				t.Errorf("newCaptchaImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCImage, tt.wantCImage) {
				t.Errorf("newCaptchaImage() = %v, want %v", gotCImage, tt.wantCImage)
			}
		})
	}
}

func TestCaptchaImageChar_drawHollowLine(t *testing.T) {
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		want    *CaptchaImageChar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.captcha.drawHollowLine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptchaImageChar.drawHollowLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptchaImageChar_drawSineLine(t *testing.T) {
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		want    *CaptchaImageChar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.captcha.drawSineLine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptchaImageChar.drawSineLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptchaImageChar_drawSlimLine(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		args    args
		want    *CaptchaImageChar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.captcha.drawSlimLine(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptchaImageChar.drawSlimLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptchaImageChar_drawBeeline(t *testing.T) {
	type args struct {
		point1    point
		point2    point
		lineColor color.RGBA
	}
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		args    args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.captcha.drawBeeline(tt.args.point1, tt.args.point2, tt.args.lineColor)
		})
	}
}

func TestCaptchaImageChar_drawNoise(t *testing.T) {
	type args struct {
		complex int
	}
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		args    args
		want    *CaptchaImageChar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.captcha.drawNoise(tt.args.complex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptchaImageChar.drawNoise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptchaImageChar_getNoiseDensityByComplex(t *testing.T) {
	type args struct {
		complex int
	}
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		args    args
		want    int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.captcha.getNoiseDensityByComplex(tt.args.complex); got != tt.want {
				t.Errorf("CaptchaImageChar.getNoiseDensityByComplex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptchaImageChar_getTextFont(t *testing.T) {
	type args struct {
		justUseFirst bool
		family       []*truetype.Font
	}
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		args    args
		want    *truetype.Font
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.captcha.getTextFont(tt.args.justUseFirst, tt.args.family); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptchaImageChar.getTextFont() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptchaImageChar_drawTextNoiseWithFontFamilySelection(t *testing.T) {
	type args struct {
		complex      int
		isSimpleFont bool
		family       []*truetype.Font
	}
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.captcha.drawTextNoiseWithFontFamilySelection(tt.args.complex, tt.args.isSimpleFont, tt.args.family); (err != nil) != tt.wantErr {
				t.Errorf("CaptchaImageChar.drawTextNoiseWithFontFamilySelection() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCaptchaImageChar_drawTextWithFontFamily(t *testing.T) {
	type args struct {
		text            string
		isSimpleFont    bool
		fontToSelection []*truetype.Font
	}
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.captcha.drawTextWithFontFamily(tt.args.text, tt.args.isSimpleFont, tt.args.fontToSelection); (err != nil) != tt.wantErr {
				t.Errorf("CaptchaImageChar.drawTextWithFontFamily() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getTextContentByMode(t *testing.T) {
	type args struct {
		config ConfigCharacter
	}
	tests := []struct {
		name               string
		args               args
		wantCaptchaContent string
		wantVerifyValue    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCaptchaContent, gotVerifyValue := getTextContentByMode(tt.args.config)
			if gotCaptchaContent != tt.wantCaptchaContent {
				t.Errorf("getTextContentByMode() gotCaptchaContent = %v, want %v", gotCaptchaContent, tt.wantCaptchaContent)
			}
			if gotVerifyValue != tt.wantVerifyValue {
				t.Errorf("getTextContentByMode() gotVerifyValue = %v, want %v", gotVerifyValue, tt.wantVerifyValue)
			}
		})
	}
}

func Test_checkConfigCharacter(t *testing.T) {
	type args struct {
		config *ConfigCharacter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkConfigCharacter(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("checkConfigCharacter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCaptchaImageChar_BinaryEncoding(t *testing.T) {
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		want    []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.captcha.BinaryEncoding(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptchaImageChar.BinaryEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptchaImageChar_WriteTo(t *testing.T) {
	tests := []struct {
		name    string
		captcha *CaptchaImageChar
		want    int64
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			got, err := tt.captcha.WriteTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("CaptchaImageChar.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CaptchaImageChar.WriteTo() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("CaptchaImageChar.WriteTo() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
