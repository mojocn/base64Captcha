// example of HTTP server that uses the captcha package.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
)

//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// base64Captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	//parse request parameters
	decoder := json.NewDecoder(r.Body)
	var param configJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	var driver base64Captcha.Driver

	//choose driver
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s,_, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

// base64Captcha verify http handler
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

	//parse request parameters
	decoder := json.NewDecoder(r.Body)
	var param configJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	//verify the captcha
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if store.Verify(param.Id, param.VerifyValue, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

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

	fmt.Println("Server is at :8777")
	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}
}
