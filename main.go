package main

import (
	"fmt"
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mws/products"
	"log"
	"net/http"
	"os"
	"time"
)

type PriceData struct {
	ASIN         string
	Amount       string
	Channel      string
	Condition    string
	ShippingTime string
	InsertTime   int64
}

func main() {
	fmt.Println()
	// API key config
	config := gmws.MwsConfig{
		SellerId:  os.Getenv("SellerId"),
		AccessKey: os.Getenv("AccessKey"),
		SecretKey: os.Getenv("SecretKey"),
		Region:    "JP",
	}

	// Create client
	productsClient, err := products.NewClient(config)
	if err != nil {
		log.Println(err)
		return
	}

	// Send request
	response := productsClient.GetLowestOfferListingsForASIN([]string{"B00JPKHFTA", "B075VQ6RFZ"})
	if response.Error != nil || response.StatusCode != http.StatusOK {
		log.Println("http Status:" + string(response.StatusCode))
		log.Println(response.Error)
		return
	}

	// responseXML to XMLNode
	xmlNode, err := gmws.GenerateXMLNode(response.Body)
	if gmws.HasErrors(xmlNode) {
		log.Println(gmws.GetErrors(xmlNode))
		return
	}

	// Get all products
	products := xmlNode.FindByKey("GetLowestOfferListingsForASINResult")
	// products to one product
	for _, product := range products {
		// Get all prices
		prices := product.FindByPath("Product.LowestOfferListings.LowestOfferListing")
		// prices to one price
		insertTime := time.Now().Unix()
		for _, price := range prices {
			temp := PriceData{
				ASIN:         product.FindByPath("Product.Identifiers.MarketplaceASIN.ASIN")[0].Value.(string),
				Amount:       price.FindByPath("Price.LandedPrice.Amount")[0].Value.(string),
				Channel:      price.FindByPath("Qualifiers.FulfillmentChannel")[0].Value.(string),
				Condition:    price.FindByPath("Qualifiers.ItemCondition")[0].Value.(string),
				ShippingTime: price.FindByPath("Qualifiers.ShippingTime.Max")[0].Value.(string),
				InsertTime:   insertTime,
			}
			fmt.Println(temp)
		}
	}
}
