package main

import (
	"flag"
	"fmt"
	"github.com/avivian/go-news-api"
)

func main() {
	apiKey := flag.String("api-key", "", "API Key for connecting to newsapi.org")
	flag.Parse()
	client := newsapi.NewClient(*apiKey, nil)

	opt := &newsapi.SourcesOptions{
		Country: newsapi.CountryGB,
	}
	sourceResp, err := client.Sources(opt)
	if err != nil {
		panic(err)
	}

	for _, source := range sourceResp.Sources {
		fmt.Println(source.Name)
	}
}
