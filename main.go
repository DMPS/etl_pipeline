package main

import (
	"fmt"
)

const url = "http://webhose.io/productFilter?token=bd9b12fb-d051-478e-8836-fa2a88849053&format=json&q=category%3Aelectronics%20-name%3AJeans&ts=1534499728377"

func saveProduct(product Product) {
	path := GeneratePath(product.Name, product.Images[0])
	err := SaveImage(product.Images[0], path)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	data, err := GetWebhoseData(url)
	if err != nil {
		fmt.Println(err)
	}
	limit := 20
	for _, product := range data.Products {
		if limit > 0 {
			saveProduct(product)
			limit--
		} else {
			break
		}
	}
}
