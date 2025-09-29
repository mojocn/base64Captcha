package main

import (
	"encoding/base64"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"io/ioutil"
	"strings"
)

func main() {
	// Create a captcha with small dimensions (120x30) as mentioned in the issue
	driver := base64Captcha.NewDriverString(30, 120, 0, 0, 4, "1234567890abcdefghjklmnpqrstuvwxyz", 
		&color.RGBA{255, 255, 255, 0}, nil, []string{})
	
	driver = driver.ConvertFonts()
	
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, b64s, answer, err := captcha.Generate()
	
	if err != nil {
		panic(err)
	}
	
	// Print info
	fmt.Printf("Small Captcha (120x30):\n")
	fmt.Printf("ID: %s\n", id)
	fmt.Printf("Answer: %s\n", answer)
	fmt.Printf("B64 length: %d\n", len(b64s))
	
	// Extract actual image data from data URL
	parts := strings.Split(b64s, ",")
	if len(parts) > 1 {
		imgData, _ := base64.StdEncoding.DecodeString(parts[1])
		ioutil.WriteFile("captcha_small.png", imgData, 0644)
		fmt.Printf("Saved small captcha to captcha_small.png\n")
	}
	
	// Create a larger captcha for comparison
	driverLarge := base64Captcha.NewDriverString(60, 240, 0, 0, 4, "1234567890abcdefghjklmnpqrstuvwxyz", 
		&color.RGBA{255, 255, 255, 0}, nil, []string{})
	
	driverLarge = driverLarge.ConvertFonts()
	
	captchaLarge := base64Captcha.NewCaptcha(driverLarge, base64Captcha.DefaultMemStore)
	idLarge, b64sLarge, answerLarge, errLarge := captchaLarge.Generate()
	
	if errLarge != nil {
		panic(errLarge)
	}
	
	// Print info
	fmt.Printf("\nLarge Captcha (240x60):\n")
	fmt.Printf("ID: %s\n", idLarge)
	fmt.Printf("Answer: %s\n", answerLarge)
	fmt.Printf("B64 length: %d\n", len(b64sLarge))
	
	// Extract actual image data from data URL
	partsLarge := strings.Split(b64sLarge, ",")
	if len(partsLarge) > 1 {
		imgDataLarge, _ := base64.StdEncoding.DecodeString(partsLarge[1])
		ioutil.WriteFile("captcha_large.png", imgDataLarge, 0644)
		fmt.Printf("Saved large captcha to captcha_large.png\n")
	}
}