package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Setup() {
	mainRouter := mux.NewRouter()
	NewMediaRoute(mainRouter)
	NewStreamlineBookRoute(mainRouter)
	http.ListenAndServe(":3060", mainRouter)
}
