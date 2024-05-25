package pixel

import (
	"testing"
)

func TestGetPixelBrightness(t *testing.T) {
	white := Pixel{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	whiteBrightness := GetPixelBrightness(&white)

	if whiteBrightness != 255 {
		t.Errorf("white pixel brightness should be 255, got %v", whiteBrightness)
	}

	black := Pixel{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}

	blackBrightness := GetPixelBrightness(&black)

	if blackBrightness != 0 {
		t.Errorf("black pixel brightness should be 0, got %v", blackBrightness)
	}

	// Add more tests for other pixel colors
}
