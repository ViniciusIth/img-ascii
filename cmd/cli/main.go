package main

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/viniciusith/img2ascii/internal/converter"
	image_util "github.com/viniciusith/img2ascii/internal/image"
	"github.com/viniciusith/img2ascii/internal/pixel"
)

func main() {
	file, _ := os.OpenFile("test.png", os.O_RDONLY, 0)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img = image_util.ScaleWidth(img, 60)

	imgMatrix := converter.Image2RGBAMatrix(img)

	saveFile, _ := os.Create("test.txt")
	for y := 0; y < len(imgMatrix); y++ {
		for x := 0; x < len(imgMatrix[y]); x++ {
			converter.GetCharacterByBrightness(pixel.GetPixelBrightness(&imgMatrix[y][x]))

			fmt.Fprintf(saveFile, "%v", converter.GetCharacterByBrightness(pixel.GetPixelBrightness(&imgMatrix[y][x])))
		}

        fmt.Fprintf(saveFile, "\n")
	}
}
