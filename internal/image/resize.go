package image_util

import (
	"image"

	"golang.org/x/image/draw"
)

// Scales the image to the given width mantaining aspect ratio
func ScaleWidth(srcImage image.Image, width int) image.Image {
	ratio := float32(srcImage.Bounds().Max.X) / float32(srcImage.Bounds().Max.Y)
	height := int(ratio * float32(width))

	return resize(srcImage, width, height, draw.ApproxBiLinear)
}

func resize(srcImage image.Image, width int, height int, scaler draw.Scaler) image.Image {
	dstRect := image.Rect(0, 0, width, height)
	dstImage := image.NewRGBA(dstRect)
	scaler.Scale(dstImage, dstRect, srcImage, srcImage.Bounds(), draw.Over, nil)
	return dstImage
}
