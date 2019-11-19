package main

import (
    "image"
    "image/color"
    "math/rand"
    "testing"

    "github.com/stretchr/testify/assert"
)

func getTestGrayValue(r uint32, g uint32, b uint32, a uint32) uint8 {
    return uint8((float64(r) * 0.21 + float64(g) * 0.72 + float64(b) * 0.07) / 256)
}

func randomUint8() uint8 {
    return uint8(rand.Uint32() % 255)
}

func TestConvertToGray(t *testing.T) {
    img := image.NewRGBA(image.Rect(0, 0, 1, 1))
    img.Set(0, 0, color.RGBA{R: randomUint8(), G: randomUint8(), B: randomUint8()})
    img.Set(0, 1, color.RGBA{R: randomUint8(), G: randomUint8(), B: randomUint8()})
    img.Set(1, 0, color.RGBA{R: randomUint8(), G: randomUint8(), B: randomUint8()})
    img.Set(1, 1, color.RGBA{R: randomUint8(), G: randomUint8(), B: randomUint8()})

    grayImage := ConvertToGray(img)

    assert.Equal(t, color.Gray{Y: getTestGrayValue(img.At(0, 0).RGBA())}, grayImage.At(0, 0))
    assert.Equal(t, color.Gray{Y: getTestGrayValue(img.At(0, 1).RGBA())}, grayImage.At(0, 1))
    assert.Equal(t, color.Gray{Y: getTestGrayValue(img.At(1, 0).RGBA())}, grayImage.At(1, 0))
    assert.Equal(t, color.Gray{Y: getTestGrayValue(img.At(1, 1).RGBA())}, grayImage.At(1, 1))
}
