package scanning

import (
	"bytes"
	"image"
	"image/png"
	"log"

	"github.com/carlogit/phash"
)

func CompareImages(img1, img2 *image.Image) (int, error) {

	hashone, err := GetImageHash(img1)
	if err != nil {
		return 0, err
	}

	hashtwo, err := GetImageHash(img2)
	if err != nil {
		return 0, err
	}

	return phash.GetDistance(hashone, hashtwo), nil
}

func GetImageHash(image *image.Image) (string, error) {

	// Create a buffer to store the encoded image
	buf := new(bytes.Buffer)

	// Encode the image as PNG to the buffer
	err := png.Encode(buf, *image)
	if err != nil {
		log.Fatal(err)
	}

	// Use the buffer as an io.Reader
	reader := bytes.NewReader(buf.Bytes())

	hash, err := phash.GetHash(reader)
	if err != nil {
		return "", err
	}

	return hash, nil
}
