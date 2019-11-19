package file

import (
    "errors"
    "fmt"
    "image"
    "image/gif"
    "image/jpeg"
    "image/png"
    "os"
    "strings"
)

func SaveNewImage(image image.Image, fileNameInput string, format string) (*os.File, error) {
    fileName, format := getFileNameAndFormat(fileNameInput, format)
    fileName = fmt.Sprintf("%s.%s", fileName, format)

    newFile, err := os.Create(fileName)
    if err != nil {
	return nil, errors.New(fmt.Sprintf("Failed to create new file: %s", err))
    }

    switch (format) {
	case "jpg":
	case "jpeg":
	    jpeg.Encode(newFile, image, nil)
	    break
	case "png":
	    png.Encode(newFile, image)
	    break
	case "gif":
	    gif.Encode(newFile, image, nil)
	    break
	default:
	    return nil, errors.New(fmt.Sprintf("Format of the new file must be gif, jpg or png. Format is %s", format))
    }

    return newFile, nil
}

func getFileNameAndFormat(fileNameInput string, format string) (string, string) {
    newFileNameParts := strings.Split(fileNameInput, ".")
    if len(newFileNameParts) < 2 {
	return fileNameInput, format
    }

    return strings.Join(newFileNameParts[:len(newFileNameParts) - 1], "."), newFileNameParts[len(newFileNameParts) - 1]
}
