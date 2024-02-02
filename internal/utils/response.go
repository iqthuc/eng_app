package utils

import (
	"net/http"
)

func ReponseData(response http.ResponseWriter, data []byte) {
	response.Header().Set("Content-Type", "application/json")
	response.Write(data)
}

func ReponseCommonError(response http.ResponseWriter) {
	response.Write([]byte("Oops! Something went wrong."))
}
