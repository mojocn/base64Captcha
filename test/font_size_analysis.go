package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Analyze current font size calculation for 30-pixel height
	height := 30
	
	rand.Seed(time.Now().UnixNano())
	
	fmt.Printf("Current font size calculation for height=%d:\n", height)
	fmt.Printf("Formula: height * (rand.Intn(7) + 7) / 16\n")
	
	for i := 0; i < 10; i++ {
		fontSize := height * (rand.Intn(7) + 7) / 16
		fmt.Printf("Iteration %d: %d * (%d + 7) / 16 = %d pixels\n", 
			i+1, height, rand.Intn(7), fontSize)
	}
	
	fmt.Printf("\nRange: %d - %d pixels\n", 
		height * 7 / 16,  // minimum
		height * 13 / 16) // maximum
		
	fmt.Printf("This gives us font sizes between %d and %d pixels for 30px height\n", 
		height * 7 / 16, height * 13 / 16)
		
	// Calculate better minimum sizes
	fmt.Printf("\nProposed improved calculation:\n")
	minFontSize := 18 // Minimum readable size
	maxVariation := 6
	
	for i := 0; i < 10; i++ {
		baseFontSize := height * 3 / 4 // Use 75% of height as base
		variation := rand.Intn(maxVariation) - maxVariation/2
		fontSize := baseFontSize + variation
		if fontSize < minFontSize {
			fontSize = minFontSize
		}
		fmt.Printf("Iteration %d: base=%d, variation=%d, final=%d pixels\n", 
			i+1, baseFontSize, variation, fontSize)
	}
}