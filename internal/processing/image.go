package processing

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"sync"

	"github.com/viniciusith/img2ascii/internal/parallel"
	"github.com/viniciusith/img2ascii/internal/pixel"
	"github.com/viniciusith/img2ascii/internal/resize"
)

func ProcessImage(filePath string, width int, adjustments func(color.RGBA) color.RGBA, outputPath string) error {
	img, err := loadImage(filePath)
	if err != nil {
		return err
	}

	if width > 0 {
		img = resize.ScaleWidth(img, width)
	}

	return saveAsASCIIArt(img, adjustments, outputPath)
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
