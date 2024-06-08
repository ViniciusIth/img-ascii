package pixel

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
