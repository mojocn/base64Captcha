package base64Captcha

const (
	imageStringDpi = 72.0
	//TxtNumbers chacters for numbers.
	TxtNumbers = "012346789"
	//TxtAlphabet characters for alphabet.
	TxtAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	//MineTypeCaptchaAudioMineType output base64 mine-type.
	MineTypeCaptchaAudioMineType = "audio/wav"
	//MineTypeCaptchaDigit output base64 mine-type.
	MineTypeCaptchaDigit = "image/png"
	//MineTypeCaptchaChar output base64 mine-type.
	MineTypeCaptchaChar = "image/png"
	//FileExtCaptchaAudio output file extension.
	FileExtCaptchaAudio = "wav"
	//FileExtCaptchaDigit output file extension.
	FileExtCaptchaDigit = "png"
	//FileExtCaptchaChar output file extension.
	FileExtCaptchaChar = "png"
)
const (
	//CaptchaComplexLower complex level lower.
	CaptchaComplexLower = iota
	//CaptchaComplexMedium complex level medium.
	CaptchaComplexMedium
	//CaptchaComplexHigh complex level high.
	CaptchaComplexHigh
)
const (
	//CaptchaModeNumber mode number.
	CaptchaModeNumber = iota
	//CaptchaModeAlphabet mode alphabet.
	CaptchaModeAlphabet
	//CaptchaModeArithmetic mode arithmetic.
	CaptchaModeArithmetic
	//CaptchaModeNumberAlphabet mode mix number and alphabet,this is also default mode.
	CaptchaModeNumberAlphabet
)

//GoTestOutputDir run go test command where the png and wav file output
var GoTestOutputDir = "/Users/ericzhou/go/src/github.com/mojocn/base64Captcha/goTestOutPutPng"

const (
	// DefaultLen Default number of digits in captcha solution.
	// 默认数字验证长度.
	DefaultLen = 6
	// MaxSkew max absolute skew factor of a single digit.
	// 图像验证码的最大干扰洗漱.
	MaxSkew = 0.7
	// DotCount Number of background circles.
	// 图像验证码干扰圆点的数量.
	DotCount = 20
)
const (
	digitFontWidth     = 11
	digitFontHeight    = 18
	digitFontBlackChar = 1
)
