package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewMediaRoute(group *mux.Router) {
	mediaDirectory := "/Users/qsoft028/Documents/english_database/AmericanStreamline/Departures"

	mediaRoute := group.PathPrefix("/media").Subrouter()

	mediaRoute.PathPrefix("/audio").Handler(http.StripPrefix("/media/audio", http.FileServer(http.Dir(mediaDirectory))))
}
