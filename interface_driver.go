package base64Captcha

// Driver captcha interface for captcha engine to to write staff
type Driver interface {
	// EncodeBinary covert to bytes
	GenerateItem(content string) (item Item, err error)
	GenerateQuestionAnswer() (q, a string)
}
