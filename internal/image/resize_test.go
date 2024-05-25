package image_util

import (
    "image"
    "testing"
)

func TestScaleWidth(t *testing.T) {
    img := image.NewNRGBA(image.Rect(0, 0, 100, 100))

    scaledImg := ScaleWidth(img, 50)

    if scaledImg.Bounds().Dx() != 50 {
        t.Errorf("Expected scaled image width to be 50, got %d", scaledImg.Bounds().Dx())
    }
}
