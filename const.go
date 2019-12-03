package base64Captcha

const (
	imageStringDpi = 72.0
	//TxtNumbers chacters for numbers.
	TxtNumbers = "012346789"
	//TxtAlphabet characters for alphabet.
	TxtAlphabet = "ABCDEFGHJKMNOQRSTUVXYZabcdefghjkmnoqrstuvxyz"
	//TxtSimpleCharaters simple numbers and alphabet
	TxtSimpleCharaters = "13467ertyiadfhjkxcvbnERTYADFGHJKXCVBN"
	//TxtChineseCharaters makes characters in chinese
	TxtChineseCharaters = "的一是在不了有和人这中大为上个国我以要他" +
		"时来用们生到作地于出就分对成会可主发年动" +
		"同工也能下过子说产种面而方后多定行学法所" +
		"民得经十三之进着等部度家电力里如水化高自" +
		"二理起小物现实加量都两体制机当使点从业本" +
		"去把性好应开它合还因由其些然前外天政四日" +
		"那社义事平形相全表间样与关各重新线内数正" +
		"心反你明看原又么利比或但质气第向道命此变" +
		"条只没结解问意建月公无系军很情者最立代想" +
		"已通并提直题党程展五果料象员革位入常文总" +
		"次品式活设及管特件长求老头基资边流路级少" +
		"图山统接知较将组见计别她手角期根论运农指" +
		"几九区强放决西被干做必战先回则任取据处队" +
		"南给色光门即保治北造百规热领七海口东导器" +
		"压志世金增争济阶油思术极交受联什认六共权" +
		"收证改清己美再采转更单风切打白教速花带安" +
		"场身车例真务具万每目至达走积示议声报斗完" +
		"类八离华名确才科张信马节话米整空元况今集" +
		"温传土许步群广石记需段研界拉林律叫且究观" +
		"越织装影算低持音众书布复容儿须际商非验连" +
		"断深难近矿千周委素技备半办青省列习响约支" +
		"般史感劳便团往酸历市克何除消构府称太准精" +
		"值号率族维划选标写存候毛亲快效斯院查江型" +
		"眼王按格养易置派层片始却专状育厂京识适属" +
		"圆包火住调满县局照参红细引听该铁价严龙飞"

	//MimeTypeCaptchaAudio output base64 mine-type.
	MimeTypeCaptchaAudio = "audio/wav"
	//MimeTypeCaptchaImage output base64 mine-type.
	MimeTypeCaptchaImage = "image/png"

	//FileExtCaptchaAudio output file extension.
	FileExtCaptchaAudio = "wav"
	//FileExtCaptchaImage output file extension.
	FileExtCaptchaImage = "png"
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
	//CaptchaModeChinese made chinese captcha .
	CaptchaModeChinese
	//CaptchaModeUseSequencedCharacters uses Sequenced Characters.
	CaptchaModeUseSequencedCharacters
)

//GoTestOutputDir run go test command where the png and wav file output

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
