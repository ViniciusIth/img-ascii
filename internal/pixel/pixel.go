package pixel

const RGBRange uint8 = 255

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func GetPixelBrightness(pixel *Pixel) uint8 {
	return uint8((0.2126 * float32(pixel.R)) + (0.7152 * float32(pixel.G)) + (0.0722 * float32(pixel.B)))
}
