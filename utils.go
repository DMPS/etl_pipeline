package main

import (
	"path/filepath"
	"strconv"
	"time"
)

func generateTimestamp() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}
func GetExtension(imagepath string) string {
	return filepath.Ext(imagepath)
}
