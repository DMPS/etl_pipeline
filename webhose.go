package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Product struct {
	URL    string `json:"url"`
	UUID   string `json:"uuid"`
	Source struct {
		SiteFull     string `json:"site_full"`
		Site         string `json:"site"`
		SiteSection  string `json:"site_section"`
		SectionTitle string `json:"section_title"`
		Country      string `json:"country"`
	} `json:"source"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	Brand            string      `json:"brand"`
	Price            float64     `json:"price"`
	Currency         string      `json:"currency"`
	OfferPrice       interface{} `json:"offer_price"`
	Model            interface{} `json:"model"`
	Manufacturer     interface{} `json:"manufacturer"`
	InStock          bool        `json:"in_stock"`
	OnSale           interface{} `json:"on_sale"`
	ProductID        string      `json:"product_id"`
	Sku              interface{} `json:"sku"`
	Mpn              interface{} `json:"mpn"`
	Colors           []string    `json:"colors"`
	AggregatedRating interface{} `json:"aggregated_rating"`
	BestRating       interface{} `json:"best_rating"`
	WorstRating      interface{} `json:"worst_rating"`
	RatingCount      interface{} `json:"rating_count"`
	ReviewsCount     interface{} `json:"reviews_count"`
	Categories       []string    `json:"categories"`
	Width            interface{} `json:"width"`
	Height           interface{} `json:"height"`
	Weight           interface{} `json:"weight"`
	Depth            interface{} `json:"depth"`
	Images           []string    `json:"images"`
	Language         string      `json:"language"`
	LastChanged      time.Time   `json:"last_changed"`
	Crawled          time.Time   `json:"crawled"`
	ProductHistory   string      `json:"product_history"`
}

type WebhoseData struct {
	Products             []Product `json:"products"`
	TotalResults         int       `json:"totalResults"`
	MoreResultsAvailable int       `json:"moreResultsAvailable"`
	Next                 string    `json:"next"`
	RequestsLeft         int       `json:"requestsLeft"`
}

func GetWebhoseData(url string) (WebhoseData, error) {
	webhoseRes, webhoseErr := http.Get(url)
	if webhoseErr != nil {
		fmt.Printf("%s", webhoseErr)
		return WebhoseData{}, webhoseErr
	}
	defer webhoseRes.Body.Close()
	contents, readErr := ioutil.ReadAll(webhoseRes.Body)
	if readErr != nil {
		fmt.Printf("%s", readErr)
		return WebhoseData{}, readErr
	}
	var data WebhoseData
	json.Unmarshal([]byte(contents), &data)
	return data, nil
}
