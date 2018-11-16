package main

import (
	"fmt"
	"sync"
)

const url = "http://webhose.io/productFilter?token=bd9b12fb-d051-478e-8836-fa2a88849053&format=json&q=category%3Aelectronics%20-name%3AJeans&ts=1534499728377"

func saveProduct(product *Product, wg *sync.WaitGroup) {
	defer wg.Done()
	name := product.Name
	imageURL := product.Images[0]
	path := GeneratePath(name, imageURL)
	err := SaveImage(imageURL, path)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	data, err := GetWebhoseData(url)
	var wg sync.WaitGroup
	if err != nil {
		fmt.Println(err)
	}
	limit := 20
	if limit > len(data.Products) {
		wg.Add(len(data.Products))
	} else {
		wg.Add(limit)
	}
	for _, product := range data.Products {
		if limit > 0 {
			go saveProduct(&product, &wg)
			limit--
		} else {
			break
		}
	}
	wg.Wait()
}
