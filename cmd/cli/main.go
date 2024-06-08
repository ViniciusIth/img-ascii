package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/viniciusith/img2ascii/internal/converter"
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

	fmt.Println(*width)

	file, _ := os.OpenFile(args[len(args)-1], os.O_RDONLY, 0)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *width >= 80 {
		img = resize.ScaleWidth(img, *width)
	}

	imgMatrix := converter.Image2RGBAMatrix(img)

	saveFile, _ := os.Create("test.txt")

	for y := 0; y < len(imgMatrix); y++ {
		for x := 0; x < len(imgMatrix[y]); x++ {
			brightness := pixel.GetPixelBrightness(&imgMatrix[y][x])
			char := converter.GetCharacterByBrightness(brightness)

			fmt.Print(char)
			fmt.Fprint(saveFile, char)
		}

		fmt.Println()
		fmt.Fprintln(saveFile)
	}

	p := pixel.Pixel{R: 100, G: 100, B: 100, A: 0}
	fmt.Println(pixel.GetPixelBrightness(&p)) // Output should be 54
}
