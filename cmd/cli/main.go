package main

import (
	"flag"
	"fmt"
	"os"

	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/viniciusith/img2ascii/internal/adjustment"
	"github.com/viniciusith/img2ascii/internal/processing"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please provide an image path")
		fmt.Println("Usage: go run cmd/cli/main.go [options] <image_path>")
		os.Exit(1)
	}

	width := flag.Int("width", 0, "Width of the image")
	contrast := flag.Float64("contrast", 0, "Adjust the contrast level")
	brightness := flag.Float64("brightness", 0, "Adjust the brightness level")
	invert := flag.Bool("invert", false, "Invert the colors")
	flag.Parse()

	imgPath := args[len(args)-1]
	err := processing.ProcessImage(imgPath, *width, createAdjustments(*contrast, *brightness, *invert), fmt.Sprintf("%s.txt", imgPath))
	if err != nil {
		fmt.Println("Error loading image:", err)
		os.Exit(1)
	}
}

func createAdjustments(contrast float64, brightness float64, invert bool) func(color.RGBA) color.RGBA {
	var adjustments []func(color.RGBA) color.RGBA

	if contrast != 0 {
		adjustments = append(adjustments, adjustment.Contrast(contrast))
	}
	if brightness != 0 {
		adjustments = append(adjustments, adjustment.Brightness(brightness))
	}
	if invert {
		adjustments = append(adjustments, adjustment.Invert())
	}

	return adjustment.Compose(adjustments...)
}
