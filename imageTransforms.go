package main

import (
	"bytes"
	"image"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
)

func ClassifyProduct(productName string) string {
	if strings.Contains(productName, "hermostat") {
		return "electric thermostat" + generateTimestamp()
	}
	if strings.Contains(productName, "rone") {
		return "drone" + generateTimestamp()
	}
	if strings.Contains(productName, "peaker") {
		return "speaker" + generateTimestamp()
	}
	if strings.Contains(productName, "adar") {
		return "radar" + generateTimestamp()
	}
	if strings.Contains(productName, "martphone") {
		return "smartphone" + generateTimestamp()
	}
	return productName
}

func GeneratePath(productName string, imagepath string) string {
	extension := GetExtension(imagepath)
	relativepath := filepath.Join("images/", ClassifyProduct(productName)+extension)
	path, _ := filepath.Abs(relativepath)
	return path
}

func TransformImage(response *http.Response) (*image.NRGBA, error) {
	buf, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return &image.NRGBA{}, readErr
	}
	src, _, decodeErr := image.Decode(bytes.NewReader(buf))
	if decodeErr != nil {
		return &image.NRGBA{}, decodeErr
	}

	image := imaging.Resize(src, 96, 96, imaging.Lanczos)

	//randomly add in blurring so that the model handles noise better
	randomNum := rand.Float64()
	if randomNum > 0.3 {
		image = imaging.Blur(image, randomNum)
	}
	return image, nil
}

func SaveImage(imageUrl string, filename string) error {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	response, getErr := client.Get(imageUrl)
	if getErr != nil {
		return getErr
	}
	defer response.Body.Close()
	image, resizeErr := TransformImage(response)
	if resizeErr != nil {
		return resizeErr
	}
	fileErr := imaging.Save(image, filename)
	if fileErr != nil {
		return fileErr
	}
	return nil
}
