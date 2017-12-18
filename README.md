# Base64 Encoding Captcha for RESTful application
[![Go Report Card](https://goreportcard.com/badge/github.com/mojocn/base64Captcha)](https://goreportcard.com/report/github.com/mojocn/base64Captcha)
[![GoDoc](https://godoc.org/github.com/mojocn/base64Captcha?status.svg)](https://godoc.org/github.com/mojocn/base64Captcha)

Package base64Captcha creates digits captcha of base64-encoding png.
base64Captcha is used for rapid development of RESTful APIs, web apps and backend services in Go.
give a string identifier to the package and it returns with a base64-encoding-png-string
#### Why Base64 for RESTful Application
      Data URIs are now supported by all major browsers. IE supports embedding images since version 8 as well.
      RESTful Application retruns small base64 image is more convenient.
A Data URI takes the format:
```
data:[<MIME-type>][;charset=<encoding>][;base64],<data>
```
CSS Image Embedding Example
```css
div.image {
  background-image:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIA...);
}
```
HTML Image Embedding Example
```html
<img alt="Embedded Image" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIA..." />
```
Similarily, if you wanted to encode an Ogg audio file, you would use the MIME type audio/ogg as follows:
```
<audio controls src="data:audio/ogg;base64,T2dnUwACAAAAAAAAAAA+..........+fm5nB6slBlZ3Fcha363d5ut7u3ni1rLoPf728l3KcK" />
```
#### Documentation

* [English](https://godoc.org/github.com/mojocn/base64Captcha)
* [中文文档](https://github.com/mojocn/base64Captcha/blob/master/README_zh.md)

#### [Playground Powered by Vuejs+elementUI+Axios](http://captcha.mojotv.cn)

[![Playground](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/captcha.png "Playground")](http://captcha.mojotv.cn/ "Playground")

## Quick Start

#### Download and Install

    go get -u github.com/mojocn/base64Captcha

#### use base64Captcha quick start a API server
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
	"strconv"
)

// base64Captcha verify http handler
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {
	//parse request parameters
	r.ParseForm()
	formData := r.Form
	captchaId := formData.Get("captchaId")
	captchaDigits := formData.Get("captchaDigits")

	//verify the captcha
	verifyResult := base64Captcha.VerifyCaptcha(captchaId, captchaDigits)

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": "error", "data": "验证失败", "msg": "captcha failed", "debug": formData}
	if verifyResult {
		body = map[string]interface{}{"code": "success", "data": "验证通过", "msg": "captcha verified", "debug": formData}
	}
	json.NewEncoder(w).Encode(body)
}

// base64Captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	//parse request parameters

	r.ParseForm()
	formData := r.Form
	captchaId := formData.Get("captchaId")
	DotCount, _ := strconv.Atoi(formData.Get("DotCount"))
	MaxSkew, _ := strconv.ParseFloat(formData.Get("MaxSkew"), 64)
	PngWidth, _ := strconv.Atoi(formData.Get("PngWidth"))
	PngHeight, _ := strconv.Atoi(formData.Get("PngHeight"))
	DefaultLen, _ := strconv.Atoi(formData.Get("DefaultLen"))

	//create base64 encoding captcha
	base64Png := base64Captcha.GenerateCaptchaPngBase64String(captchaId, PngWidth, PngHeight, DotCount, DefaultLen, MaxSkew)
	//or you can do this
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": 1, "data": base64Png, "msg": "success", "debug": formData}
	json.NewEncoder(w).Encode(body)
}

//start a net/http server
func main() {
	//serve Vuejs+ElementUI+Axios Web Application
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//api for create captcha
	http.HandleFunc("/api/getCaptcha", generateCaptchaHandler)
	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)

	fmt.Println("Server is at localhost:777")
	if err := http.ListenAndServe(":777", nil); err != nil {
		log.Fatal(err)
	}
}
```
#### base64Captcha package function
-  `func GenerateCaptchaPngBase64String(identifier string, pngWidth, pngHeight, DotCount, digitsLen int, maxSkew float64) string` return base64-png-captcha
-  `func GenerateCaptchaPngBase64StringDefault(identifier string) string ` default settings width=240 height=70 dot-count=20 digits-len=6 skew-factor=0.7
-  `func VerifyCaptcha(identifier, digits string) bool` verify the captcha-png-numbers by identifierKey
-  `func RandomId() string` Server Create Random IdentifierKey

#### Build and Run the Demo
    cd $GOPATH/src/github.com/mojocn/captcha/examples
    go run main.go

#### demo nginx configuration `captcha.mojotv.cn.config`
```
server {
        listen 80;
        server_name captcha.mojotv.cn;
        charset utf-8;

        location / {
            try_files /_not_exists_ @backend;
        }
        location @backend {
           proxy_set_header X-Forwarded-For $remote_addr;
           proxy_set_header Host $http_host;
           proxy_pass http://127.0.0.1:777;
        }
        access_log  /home/wwwlogs/captcha.mojotv.cn.log;
}
```
#### Go to [http://localhost:777](http://localhost:777)

Congratulations! You've just built your first **base64Captcha-APIs** app.
Any question you can leave a message. If you like the package please star this repo
## License

base64Captcha source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).
