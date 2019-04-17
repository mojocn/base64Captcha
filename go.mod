module github.com/mojocn/base64Captcha

go 1.12

replace (
	golang.org/x/image => github.com/golang/image v0.0.0-20190417020941-4e30a6eb7d9a
	golang.org/x/text => github.com/golang/text v0.3.0

)

require (
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	golang.org/x/image v0.0.0-00010101000000-000000000000
)
