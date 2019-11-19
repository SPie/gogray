package main

import (
    "fmt"
    "os"

    "github.com/spie/gogray/file"
)

func main() {
    if len(os.Args) < 3 {
	fmt.Println("gogray FILENAME NEWFILENAME")
	os.Exit(1)
    }

    imageFile, err := os.Open(os.Args[1])
    if err != nil {
	fmt.Printf("Failed to open the file: %s", err)
	os.Exit(1)
    }
    defer imageFile.Close()

    img, format, err := OpenImage(imageFile)
    if err != nil {
	fmt.Printf("Failed to load the image: %s", err)
	os.Exit(1)
    }

    grayImage := ConvertToGray(img)

    newFile, err := file.SaveNewImage(grayImage, os.Args[2], format)
    if err != nil {
	fmt.Println(err)
	os.Exit(1)
    }
    defer newFile.Close()

    fmt.Printf("Created new gray image: %s", newFile.Name())
}
