package database

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
)

var EsClient *elasticsearch.Client

func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ES_HOST"),
		},
	}

	var err error
	EsClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	// Test the connection
	res, err := EsClient.Info()
	if err != nil {
		log.Fatalf("Error getting response from Elasticsearch: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}
