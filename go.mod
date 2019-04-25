module github.com/mojocn/base64Captcha

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190424203555-c05e17bb3b2d
	golang.org/x/image => github.com/golang/image v0.0.0-20190424155947-59b11bec70c7
	golang.org/x/net => github.com/golang/net v0.0.0-20190424112056-4829fb13d2c6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190425045458-9f0b1ff7b46a
	golang.org/x/text => github.com/golang/text v0.3.1
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190425001055-9e44c1c40307
)

require (
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	golang.org/x/image v0.0.0-20190424155947-59b11bec70c7
)
