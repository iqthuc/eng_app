package controller

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func RequestMedia(w http.ResponseWriter, r *http.Request) {
	fileName := path.Base(r.URL.Path)

	// mediaDirectory := "/home/iqthuc/Downloads/data compiled/resources/assets/AmericanStreamline/" // for my linux
	mediaDirectory := "/Users/qsoft028/Documents/english_database/AmericanStreamline/Departures" // for my mac

	files, err := os.ReadDir(mediaDirectory)

	if err != nil {
		http.NotFound(w, r)
	} else {
		for _, file := range files {
			if file.Name() == fileName {
				http.ServeFile(w, r, filepath.Join(mediaDirectory, fileName))
				return
			}
		}
		http.NotFound(w, r)
	}
}
