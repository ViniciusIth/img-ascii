package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"sync"

	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/viniciusith/img2ascii/internal/adjustment"
	"github.com/viniciusith/img2ascii/internal/parallel"
	"github.com/viniciusith/img2ascii/internal/pixel"
	"github.com/viniciusith/img2ascii/internal/resize"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please provide an image path")
		fmt.Println("Usage: go run cmd/cli/main.go [options] <image_path>")
		os.Exit(1)
	}

	width := flag.Int("width", 0, "Width of the image")
	flag.Parse()

	imgPath := args[len(args)-1]
	img, err := loadImage(imgPath)
	if err != nil {
		fmt.Println("Error loading image:", err)
		os.Exit(1)
	}

	if *width > 0 {
		img = resize.ScaleWidth(img, *width)
	}

	adjustments := createAdjustments()

	if err := saveAsASCIIArt(img, adjustments, "test.txt"); err != nil {
		fmt.Println("Error saving ASCII art:", err)
		os.Exit(1)
	}
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func createAdjustments() func(color.RGBA) color.RGBA {
	contrast := adjustment.Contrast(3)
	brightness := adjustment.Brightness(0.9)
	invert := adjustment.Invert()
	return adjustment.Compose(contrast, brightness, invert)
}

func saveAsASCIIArt(img image.Image, adjustments func(color.RGBA) color.RGBA, outputPath string) error {
	saveFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer saveFile.Close()

	var wg sync.WaitGroup
	channel := make(chan []string, img.Bounds().Max.Y)

	parallel.ProcessInParallel(img.Bounds().Max.Y, func(start, end int) {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			pixel.ProcessImageSegment(pixel.Task{StartY: start, EndY: end, Img: img, Result: channel}, adjustments)
		}(start, end)
	})

	wg.Wait()
	close(channel)

	for rows := range channel {
		for _, row := range rows {
			fmt.Fprintln(saveFile, row)
		}
	}

	return nil
}
