package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("banners/standard")
	lines := strings.Split(string(data), "\n")
	
	// Print H character (ASCII 72, index 40)
	fmt.Println("Character H (index 40):")
	start := 40 * 9
	for i := 0; i < 8; i++ {
		fmt.Printf("'%s'\n", lines[start+i])
	}
	
	fmt.Println("\nCharacter e (ASCII 101, index 69):")
	start = 69 * 9
	for i := 0; i < 8; i++ {
		fmt.Printf("'%s'\n", lines[start+i])
	}
}