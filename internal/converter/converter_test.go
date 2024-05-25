package converter

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"
)

func TestImage2RGBAMatrix(t *testing.T) {
	file, _ := os.OpenFile("test.jpg", os.O_RDONLY, 0)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	imgMatrix := Image2RGBAMatrix(img)

	if len(imgMatrix) != 320 {
		t.Errorf("expected 320 columns, got %v", len(imgMatrix))
	}

	if len(imgMatrix[0]) != 400 {
		t.Errorf("expected 400 rows, got %v", len(imgMatrix[0]))
	}
}
