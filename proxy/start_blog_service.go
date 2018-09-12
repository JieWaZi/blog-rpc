package main

import "github.com/blog/api"

func main() {
	firstCalculationService := api.BlogService{}
	firstCalculationService.Start()
}