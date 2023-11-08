package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Setup() {
	mainRouter := mux.NewRouter()
	NewMediaRoute(mainRouter)
	http.ListenAndServe(":3060", mainRouter)
}
