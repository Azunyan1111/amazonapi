package main

import (
	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mws/products"
	"log"
	"os"
	"fmt"
	"net/http"
)


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
	if err != nil{
		log.Println(err)
		return
	}

	// Send request
	response := productsClient.GetLowestOfferListingsForASIN([]string{"B00JPKHFTA","B075VQ6RFZ"})
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

	// XML Parse
	xmlNode.PrintXML()

	// Get all products
	products := xmlNode.FindByKey("GetLowestOfferListingsForASINResult")
	// products to one product
	for _, product := range products{
		// ASIN
		fmt.Println(product.FindByPath("Product.Identifiers.MarketplaceASIN.ASIN")[0].Value)
		// Get all prices
		prices := product.FindByPath("Product.LowestOfferListings.LowestOfferListing.Price.LandedPrice.Amount")
		// prices to one price
		for _, price := range prices{
			fmt.Println(price.Value)
		}
	}
}

