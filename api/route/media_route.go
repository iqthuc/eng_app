package route

import (
	controller "eng_app_module/api/controller/media"
	"net/http"
)

func NewMediaRoute(server *http.ServeMux) {
	audioRoutePrefix := "/media/"

	server.HandleFunc(audioRoutePrefix, controller.RequestMedia)

}

// func NewMediaRoute(group *mux.Router) {
// 	mediaDirectory := "/home/iqthuc/Downloads/data compiled/resources/assets/AmericanStreamline/"

// 	mediaRoute := group.PathPrefix("/media").Subrouter()

// 	mediaRoute.PathPrefix("/audio/{filepath:.+}").Handler(http.StripPrefix("/media/audio", http.FileServer(http.Dir(mediaDirectory))))
// }
