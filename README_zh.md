# base64Captcha快速生成base64编码图片验证码字符串
Base64是网络上最常见的用于传输8Bit字节代码的编码方式之一。Base64编码可用于在HTTP环境下传递较长的标识信息, 直接把base64当成是字符串方式的数据就好了
减少了http请求；数据就是图片；
为APIs微服务而设计
#### 为什么base64图片 for RESTful 服务
      Data URIs 支持大部分浏览器,IE8之后也支持.
      小图片使用base64响应对于RESTful服务来说更便捷
CSS Image 嵌入base64图片
```css
div.image {
  background-image:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIA...);
}
```
HTML 嵌入base64图片
```html
<img alt="Embedded Image" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADIA..." />
```
#### [godoc文档](https://godoc.org/github.com/mojocn/base64Captcha)

#### 在线Demo [Playground Powered by Vuejs+elementUI+Axios](http://captcha.mojotv.cn)

[![Playground](https://raw.githubusercontent.com/mojocn/base64Captcha/master/examples/static/captcha.png "Playground")](http://captcha.mojotv.cn/ "Playground")


## 快速开始

#### 安装golang包

    go get -u github.com/mojocn/base64Captcha

#### 使用golang搭建API服务
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
	//接收客户端发送来的请求参数
	r.ParseForm()
	formData := r.Form
	captchaId := formData.Get("captchaId")
	captchaDigits := formData.Get("captchaDigits")

	//比较图像验证码
	verifyResult := base64Captcha.VerifyCaptcha(captchaId, captchaDigits)

	//设置json响应
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": "error", "data": "验证失败", "msg": "captcha failed", "debug": formData}
	if verifyResult {
		body = map[string]interface{}{"code": "success", "data": "验证通过", "msg": "captcha verified", "debug": formData}
	}
	json.NewEncoder(w).Encode(body)
}

// base64Captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	//接收客户端发送来的请求参数

	r.ParseForm()
	formData := r.Form
	captchaId := formData.Get("captchaId")
	DotCount, _ := strconv.Atoi(formData.Get("DotCount"))
	MaxSkew, _ := strconv.ParseFloat(formData.Get("MaxSkew"), 64)
	PngWidth, _ := strconv.Atoi(formData.Get("PngWidth"))
	PngHeight, _ := strconv.Atoi(formData.Get("PngHeight"))
	DefaultLen, _ := strconv.Atoi(formData.Get("DefaultLen"))

	//创建base64图像验证码
	base64Png := base64Captcha.GenerateCaptchaPngBase64String(captchaId, PngWidth, PngHeight, DotCount, DefaultLen, MaxSkew)
	//你也可以是用默认参数 生成图像验证码
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	//设置json响应

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": 1, "data": base64Png, "msg": "success", "debug": formData}
	json.NewEncoder(w).Encode(body)
}

//启动golang net/http 服务器
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
#### base64Captcha包主要方法
-  自定参数返回验证码base64png
`func GenerateCaptchaPngBase64String(identifier string, pngWidth, pngHeight, DotCount, digitsLen int, maxSkew float64) string`
- default settings width=240 height=70 dot-count=20 digits-len=6 skew-factor=0.7
`func GenerateCaptchaPngBase64StringDefault(identifier string) string `

- 使用默认设置生成验证码base64png
`func VerifyCaptcha(identifier, digits string) bool`
- 参数随机的idKey
`func RandomId() string`
#### 运行demo代码
    cd $GOPATH/src/github.com/mojocn/captcha/examples
    go run main.go
#### golang-demo nginx 配置 `captcha.mojotv.cn.config`
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
#### 访问 [http://localhost:777](http://localhost:777)

如果喜欢,请star 非常感谢.

## License

base64Captcha source code is licensed under the Apache Licence, Version 2.0
(http://www.apache.org/licenses/LICENSE-2.0.html).
