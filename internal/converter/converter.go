package converter

import (
	"image"

	"github.com/viniciusith/img2ascii/internal/pixel"
)

func Image2RGBAMatrix(img image.Image) [][]pixel.Pixel {
	size := img.Bounds()
	imgHeight := size.Max.Y
	imgWidth := size.Max.X

	// Initialize the matrix with the appropriate dimensions
	imgMatrix := make([][]pixel.Pixel, imgHeight)
	for i := range imgMatrix {
		imgMatrix[i] = make([]pixel.Pixel, imgWidth)
	}

	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			pixelAt := img.At(x, y)
			R, G, B, A := pixelAt.RGBA()

			rgbaPixel := pixel.Pixel{
				R: uint8(R),
				G: uint8(G),
				B: uint8(B),
				A: uint8(A),
			}
			imgMatrix[y][x] = rgbaPixel
		}
	}

	return imgMatrix
}

func GetCharacterByBrightness(brightness uint8) string {
	var CHARS = []rune{' ', '_', '.', ',', '-', '=', '+', ':', ';', 'c', 'b', 'a', '?', '!', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '$', 'W', '#', '@', 'Ñ'}

	return string(CHARS[brightness/uint8(len(CHARS))])
}