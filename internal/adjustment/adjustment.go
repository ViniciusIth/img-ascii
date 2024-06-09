package adjustment

import (
	"image/color"
)

// Contrast returns a function that adjusts the contrast of the image.
func Contrast(change float64) func(color.RGBA) color.RGBA {
	lookup := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		contrast := uint8(clamp(((((float64(i)/255.0)-0.5)*(1+change))+0.5)*255, 0, 255))
		lookup[i] = uint8(contrast)
	}

	return func(c color.RGBA) color.RGBA {
		return color.RGBA{
			R: lookup[c.R],
			G: lookup[c.G],
			B: lookup[c.B],
			A: c.A,
		}
	}
}

// Invert returns a function that inverts the colors.
func Invert() func(color.RGBA) color.RGBA {
	return func(c color.RGBA) color.RGBA {
		return color.RGBA{
			R: 255 - c.R,
			G: 255 - c.G,
			B: 255 - c.B,
			A: c.A,
		}
	}
}

// Brightness returns a function that adjusts the brightness of the image.
func Brightness(change float64) func(color.RGBA) color.RGBA {
	return func(c color.RGBA) color.RGBA {
		adjust := func(val uint8) uint8 {
			return uint8(clamp(float64(val)+change*255, 0, 255))
		}
		return color.RGBA{
			R: adjust(c.R),
			G: adjust(c.G),
			B: adjust(c.B),
			A: c.A,
		}
	}
}

// Compose applies multiple color adjustments sequentially
func Compose(adjustments ...func(color.RGBA) color.RGBA) func(color.RGBA) color.RGBA {
	return func(c color.RGBA) color.RGBA {
		for _, adjustment := range adjustments {
			c = adjustment(c)
		}

		return c
	}
}

// clamps a value between a minimum and maximum.
func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
