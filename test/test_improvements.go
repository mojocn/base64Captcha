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
	fmt.Println("Testing improved captcha for 120x30 dimensions...")
	
	// Create a captcha with small dimensions (120x30) as mentioned in the issue
	// Using the new parameters for better clarity
	driver := base64Captcha.NewDriverString(30, 120, 0, 0, 4, "1234567890abcdefghjklmnpqrstuvwxyz", 
		&color.RGBA{255, 255, 255, 0}, nil, []string{})
	
	// Enable bold text for better clarity
	driver.Bold = true
	
	// Let's also set custom font sizes
	driver.MinFontSize = 18  // Ensure minimum readable size
	driver.MaxFontSize = 24  // Maximum for variation
	
	driver = driver.ConvertFonts()
	
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	
	// Generate multiple samples to see consistency
	for i := 0; i < 5; i++ {
		id, b64s, answer, err := captcha.Generate()
		
		if err != nil {
			panic(err)
		}
		
		fmt.Printf("\nImproved Captcha #%d (120x30, bold, font size 18-24):\n", i+1)
		fmt.Printf("ID: %s\n", id)
		fmt.Printf("Answer: %s\n", answer)
		fmt.Printf("B64 length: %d\n", len(b64s))
		
		// Extract actual image data from data URL
		parts := strings.Split(b64s, ",")
		if len(parts) > 1 {
			imgData, _ := base64.StdEncoding.DecodeString(parts[1])
			filename := fmt.Sprintf("captcha_improved_%d.png", i+1)
			ioutil.WriteFile(filename, imgData, 0644)
			fmt.Printf("Saved to %s\n", filename)
		}
	}
	
	// Test without bold for comparison
	fmt.Println("\n--- Testing without bold for comparison ---")
	driver.Bold = false
	
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("\nNon-bold Captcha (120x30, font size 18-24):\n")
	fmt.Printf("ID: %s\n", id)
	fmt.Printf("Answer: %s\n", answer)
	
	parts := strings.Split(b64s, ",")
	if len(parts) > 1 {
		imgData, _ := base64.StdEncoding.DecodeString(parts[1])
		ioutil.WriteFile("captcha_non_bold.png", imgData, 0644)
		fmt.Printf("Saved to captcha_non_bold.png\n")
	}
}