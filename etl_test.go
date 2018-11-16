package main

import "testing"

func TestGetExtension(t *testing.T) {
	result := GetExtension("https://i5.walmartimages.com/asr/80beb8e1-e5f9-4156-be86-d136cf59cb5c_1.04d884e2cafb42c46261fce9524c47da.jpeg")
	expected := ".jpeg"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func TestGeneratePath(t *testing.T) {
	result := GeneratePath("Nest Learning Touch Home/Office Heating/Cooling Smart Thermostat with WiFi", "https://i5.walmartimages.com/asr/80beb8e1-e5f9-4156-be86-d136cf59cb5c_1.04d884e2cafb42c46261fce9524c47da.jpeg")
	expected := "/Users/gasteig/Documents/Code/go/src/github.com/dmps/imageClassifier/images/electric thermostat" + generateTimestamp() + ".jpeg"

	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}
