package route

import (
	controller "eng_app_module/api/controller/streamline"

	"github.com/gorilla/mux"
)

func NewStreamlineBookRoute(group *mux.Router) {
	streamlineBookRoute := group.PathPrefix("/streamline").Subrouter()
	streamlineBookRoute.Path("/lessons").HandlerFunc(controller.GetLessons).Methods("GET")
	streamlineBookRoute.Path("/lessons/{id}").HandlerFunc(controller.GetLessonDetail).Methods("GET")
	streamlineBookRoute.Path("/books").HandlerFunc(controller.GetBooks).Methods("GET")
}
