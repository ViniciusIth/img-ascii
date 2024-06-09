package pixel

import (
	"image"
	"image/color"
)

const RGBRange uint8 = 255

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func GetPixelBrightness(pixel *Pixel) uint8 {
	// Calculate brightness using one of the formulas for luminance
	brightness := uint8((0.2126 * float32(pixel.R)) + (0.7152 * float32(pixel.G)) + (0.0722 * float32(pixel.B)))

	// Adjust brightness based on alpha value
	adjustedBrightness := float32(brightness) * (float32(pixel.A) / 255.0)

	// Should I check if the brightness is greater than 255 or smaller than 0?
	// I can't really see them being trespassed
	return uint8(adjustedBrightness)
}

func GetCharacterByBrightness(brightness uint8) string {
	var CHARS = []rune{' ', '_', '.', ',', '-', '=', '+', ':', ';', 'c', 'b', 'a', '?', '!', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '$', 'W', '#', '@', 'Ñ'}

	return string(CHARS[int(brightness)*(len(CHARS)-1)/255])
}

type Task struct {
	StartY int             // Starting row index for this task
	EndY   int             // Ending row index for this task
	Img    image.Image     // The image being processed
	Result chan<- []string // Channel to send the result (rows of ASCII art)
}

func ProcessImageSegment(task Task, adjustments func(color.RGBA) color.RGBA) {
	rows := make([]string, task.EndY-task.StartY)
	for y := task.StartY; y < task.EndY; y++ {
		row := ""
		for x := 0; x < task.Img.Bounds().Max.X; x++ {
			pixelAt := task.Img.At(x, y)
			r, g, b, a := pixelAt.RGBA()
			RGBAPixel := Pixel{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}

			adjustedPixel := adjustments(color.RGBA{RGBAPixel.R, RGBAPixel.G, RGBAPixel.B, RGBAPixel.A})

			brightness := GetPixelBrightness(&Pixel{R: adjustedPixel.R, G: adjustedPixel.G, B: adjustedPixel.B, A: adjustedPixel.A})
			char := GetCharacterByBrightness(brightness)
			row += char
		}
		rows[y-task.StartY] = row
	}
	task.Result <- rows
}
