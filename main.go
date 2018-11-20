package main

import (
	"fmt"
	"strconv"
	"sync"
)

const baseurl = "http://webhose.io/productFilter?token=bd9b12fb-d051-478e-8836-fa2a88849053&format=json&q=category%3Aelectronics%20-name%3AJeans&ts=1534499728377"

func saveProduct(productName, imageURL string, wg *sync.WaitGroup) {
	defer wg.Done()
	path := GeneratePath(productName, imageURL)
	err := SaveImage(imageURL, path)
	if err != nil {
		fmt.Println(err)
	}
}

func savePage(url string, wg *sync.WaitGroup) {
	page, err := GetWebhoseData(url)
	fmt.Println("Page " + strconv.Itoa(MAX_PAGES))
	if err != nil {
		fmt.Println(err)
	} else {
		wg.Add(len(page.Products))
		for _, product := range page.Products {
			name := product.Name
			images := product.Images
			if len(images) > 0 {
				go saveProduct(name, images[0], wg)
			}
		}
		if MAX_PAGES > 0 && page.Next != "" {
			MAX_PAGES--
			savePage(page.Next, wg)
		}
	}
}

var MAX_PAGES = 10 //limits the number of pages the pipeline can ask for

func main() {
	var wg sync.WaitGroup
	savePage(baseurl, &wg)
	wg.Wait()
}
