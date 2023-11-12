package route

import (
	"eng_app_module/api/controller"

	"github.com/gorilla/mux"
)

func NewStreamlineBookRoute(group *mux.Router) {
	streamlineBookRoute := group.PathPrefix("/streamline").Subrouter()
	streamlineBookRoute.PathPrefix("/books").HandlerFunc(controller.GetBooks).Methods("GET")
	streamlineBookRoute.PathPrefix("/lessons").HandlerFunc(controller.GetLessons).Methods("GET")
	streamlineBookRoute.PathPrefix("/lessons/{id}").HandlerFunc(controller.GetBooks).Methods("GET")
}
