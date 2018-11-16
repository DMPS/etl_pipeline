package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ClassifyProduct(productName string) string {
	if strings.Contains(productName, "hermostat") {
		return "electric thermostat" + generateTimestamp()
	}
	if strings.Contains(productName, "rone") {
		return "drone" + generateTimestamp()
	}
	return ""
}

func GeneratePath(productName string, imagepath string) string {
	extension := GetExtension(imagepath)
	relativepath := filepath.Join("images/", ClassifyProduct(productName)+extension)
	path, _ := filepath.Abs(relativepath)
	return path
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
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()
	_, fileErr = io.Copy(file, response.Body)
	if fileErr != nil {
		return fileErr
	}
	return nil
}
