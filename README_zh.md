# 快速生成base64编码图片验证码字符串.base64图形验证码(captcha)为golang而设计.
支持多种样式,算术,数字,字母,混合模式,语音模式.

Base64是网络上最常见的用于传输8Bit字节代码的编码方式之一。Base64编码可用于在HTTP环境下传递较长的标识信息, 直接把base64当成是字符串方式的数据就好了
减少了http请求；数据就是图片；
为APIs微服务而设计
#### 为什么base64图片 for RESTful 服务
      Data URIs 支持大部分浏览器,IE8之后也支持.
      小图片使用base64响应对于RESTful服务来说更便捷

#### [godoc文档](https://godoc.org/github.com/mojocn/base64Captcha)

#### 在线Demo [Playground Powered by Vuejs+elementUI+Axios](http://captcha.mojotv.cn)


[![Playground](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/captcha.png "Playground")](http://captcha.mojotv.cn/ "Playground")
[![28+58.png](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/28%2B58%3D%3F.png)](http://captcha.mojotv.cn/ "Playground")
[![ACNRfd.png](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/ACNRfd.png)](http://captcha.mojotv.cn/ "Playground")
[![rW4npZ.png](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/rW4npZ.png)](http://captcha.mojotv.cn/ "Playground")
[![99+73.png](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/99%2B73%3D%3F.png)](http://captcha.mojotv.cn/ "Playground")
[![ctOv6N.png](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/ctOv6N.png)](http://captcha.mojotv.cn/ "Playground")
[![gGncJC.png](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/gGncJC.png)](http://captcha.mojotv.cn/ "Playground")
[![108360.png](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/108360.png)](http://captcha.mojotv.cn/ "Playground")
[wav file](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/1lNMVxfysfSQJXvjR1LX.wav)


### 安装golang包

    go get -u github.com/mojocn/base64Captcha

对于中国大陆Gopher `go get golang.org/x/image` 失败解决方案:
#### 方法1
```bash
    mkdir -p $GOPATH/src/golang.org/x
    cd $GOPATH/src/golang.org/x
    git clone https://github.com/golang/image.git
```
#### 方法2
- go version > 1.11
- set env `GOPROXY=https://goproxy.io`
- ![](_examples/static/gomodproxy.png)

###  创建图像验证码
```
import "github.com/mojocn/base64Captcha"
func demoCodeCaptchaCreate() {
	//config struct for digits
	//数字验证码配置
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	//config struct for audio
	//声音验证码配置
	var configA = base64Captcha.ConfigAudio{
		CaptchaLen: 6,
		Language:   "zh",
	}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height:             60,
		Width:              240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
    
	fmt.Println(idKeyA, base64stringA, "\n")
	fmt.Println(idKeyC, base64stringC, "\n")
	fmt.Println(idKeyD, base64stringD, "\n")
}

```
### 验证图像验证码
```
import "github.com/mojocn/base64Captcha"
func verfiyCaptcha(idkey,verifyValue string){
    verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
    if verifyResult {
        //success
    } else {
        //fail
    }
}
```
#### 使用golang搭建API服务
```go
// example of HTTP server that uses the captcha package.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
)

//ConfigJsonBody json request body.
type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

// base64Captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	//parse request parameters
	//接收客户端发送来的请求参数
	decoder := json.NewDecoder(r.Body)
	var postParameters ConfigJsonBody
	err := decoder.Decode(&postParameters)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	//create base64 encoding captcha
	//创建base64图像验证码

	var config interface{}
	switch postParameters.CaptchaType {
	case "audio":
		config = postParameters.ConfigAudio
	case "character":
		config = postParameters.ConfigCharacter
	default:
		config = postParameters.ConfigDigit
	}
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha(postParameters.Id, config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	//or you can do this
	//你也可以是用默认参数 生成图像验证码
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	//set json response
	//设置json响应

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": 1, "data": base64Png, "captchaId": captchaId, "msg": "success"}
	json.NewEncoder(w).Encode(body)
}
// base64Captcha verify http handler
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

	//parse request parameters
	//接收客户端发送来的请求参数
	decoder := json.NewDecoder(r.Body)
	var postParameters ConfigJsonBody
	err := decoder.Decode(&postParameters)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	//verify the captcha
	//比较图像验证码
	verifyResult := base64Captcha.VerifyCaptcha(postParameters.Id, postParameters.VerifyValue)

	//set json response
	//设置json响应
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": "error", "data": "验证失败", "msg": "captcha failed"}
	if verifyResult {
		body = map[string]interface{}{"code": "success", "data": "验证通过", "msg": "captcha verified"}
	}
	json.NewEncoder(w).Encode(body)
}

//start a net/http server
//启动golang net/http 服务器
func main() {

	//serve Vuejs+ElementUI+Axios Web Application
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//api for create captcha
	//创建图像验证码api
	http.HandleFunc("/api/getCaptcha", generateCaptchaHandler)

	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)

	fmt.Println("Server is at localhost:3333")
	if err := http.ListenAndServe("localhost:3333", nil); err != nil {
		log.Fatal(err)
	}
}
```
#### [使用redis做储存](examples_redis/main.go)
#### 运行demo代码
    cd $GOPATH/src/github.com/mojocn/captcha/_examples
    go run main.go

#### 访问 `http://localhost:777`


## License

base64Captcha source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).
