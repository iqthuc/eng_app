package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Setup() {
	mainRouter := mux.NewRouter()
	NewMediaRoute(mainRouter)
	NewStreamlineBookRoute(mainRouter)
	http.ListenAndServe(":3060", mainRouter)
}
