package routes

import (
	"github.com/gorilla/mux"
	"github.com/sir-geronimo/go-workshop/pkg/controllers"
)

var RegisterBookStoreRouters = func(r *mux.Router) {
	r.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
