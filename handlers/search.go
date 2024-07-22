package handlers

import (
	"bookshelf-api/database"
	"bookshelf-api/models"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func IndexBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Indexing book: %+v", book)

	body, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := database.EsClient.Index(
		"books",
		bytes.NewReader(body),
		database.EsClient.Index.WithDocumentID(strconv.Itoa(int(book.ID))),
		database.EsClient.Index.WithRefresh("true"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}
		log.Fatalf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": "book indexed"})
}

func SearchBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]

	var buf bytes.Buffer
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"title", "author"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := database.EsClient.Search(
		database.EsClient.Search.WithContext(context.Background()),
		database.EsClient.Search.WithIndex("books"),
		database.EsClient.Search.WithBody(&buf),
		database.EsClient.Search.WithTrackTotalHits(true),
		database.EsClient.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}
		log.Fatalf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
	}

	var esResponse models.EsSearchResponse
	if err := json.NewDecoder(res.Body).Decode(&esResponse); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	books := make([]models.Book, len(esResponse.Hits.Hits))
	for i, hit := range esResponse.Hits.Hits {
		books[i] = hit.Source
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
