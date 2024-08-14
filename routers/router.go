package routers

import (
	"bookshelf-api/handlers"
	"bookshelf-api/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	bookHandler := handlers.NewBookHandler()

	router := mux.NewRouter()
	router.Use(middlewares.JSONMiddleware)

	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/refresh", handlers.Refresh).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)
	api.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	api.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	api.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods("GET")
	api.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")

	api.HandleFunc("/books/index", handlers.IndexBook).Methods("POST")
	api.HandleFunc("/books/search/{query}", handlers.SearchBooks).Methods("GET")

	return router
}
