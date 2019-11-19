package main

import (
    "image"
    "image/color"
    "os"
)

func OpenImage(imageFile *os.File) (image.Image, string, error) {
    return image.Decode(imageFile)
}

func ConvertToGray(img image.Image) *image.Gray {
    bounds := img.Bounds()
    grayImage := image.NewGray(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))
    
    for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
	    r, g, b, _ := img.At(x, y).RGBA()
	    grayImage.Set(x, y, color.Gray{Y: getGrayValue(r, g, b)})
	}
    }

    return grayImage
}

func getGrayValue(r uint32, g uint32, b uint32) uint8 {
    return uint8((float64(r) * 0.21 + float64(g) * 0.72 + float64(b) * 0.07) / 256)
}
