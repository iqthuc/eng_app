package route

import (
	controller "eng_app_module/api/controller/media"
	"net/http"

	"github.com/gorilla/mux"
)

func NewMediaRoute(group *mux.Router) {
	mediaDirectory := "/home/iqthuc/Downloads/data compiled/resources/assets/AmericanStreamline/"

	mediaRoute := group.PathPrefix("/media").Subrouter()
	mediaRoute.Path("/drive").HandlerFunc(controller.GetGoogleSheetData)
	mediaRoute.PathPrefix("/audio/{filepath:.+}").Handler(http.StripPrefix("/media/audio", http.FileServer(http.Dir(mediaDirectory))))
}
