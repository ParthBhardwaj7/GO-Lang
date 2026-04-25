package routes

import (
	"github.com/gorilla/mux"
	"github.com/parthbhardwaj93/Bookstore-Management-APIs/package/controllers"

)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookid}",controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookid}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}",controllers.DeleteBook).Methods("DELETE")
}