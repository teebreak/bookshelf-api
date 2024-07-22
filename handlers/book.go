package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
	"time"

	"bookshelf-api/database"
	"bookshelf-api/models"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	database.DB.Create(&book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	database.DB.Find(&books)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Check Redis cache
	val, err := database.RDB.Get(database.Ctx, id).Result()
	if errors.Is(err, redis.Nil) {
		// Key does not exist in Redis, fetch from DB
		var book models.Book
		if err := database.DB.First(&book, id).Error; err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Cache the result in Redis
		jsonBook, _ := json.Marshal(book)
		database.RDB.Set(database.Ctx, id, jsonBook, 10*time.Minute)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		// Key exists in Redis, return the cached value
		var book models.Book
		json.Unmarshal([]byte(val), &book)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var book models.Book
	database.DB.First(&book, id)
	json.NewDecoder(r.Body).Decode(&book)
	database.DB.Save(&book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var book models.Book
	database.DB.Delete(&book, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("{success: true}")
}
