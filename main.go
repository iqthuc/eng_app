package main

import (

	// "log"
	database "eng_app_module/config"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func getSomeData(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConfigDatabase()
	if err != nil {
		fmt.Println("The database connection error:", err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var title string
		var discard any
		err := rows.Scan(&id, &title, &discard, &discard)
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(fmt.Sprintf("id: %s---title:%s \n", id, title)))
	}
}
func testStreamMp3(w http.ResponseWriter, r *http.Request) {
	// audioFile, err := os.Open("/home/iqthuc/Downloads/Nang-Tho-Hoang-Dung.mp3")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer audioFile.Close()
	// w.Header().Set("Content-Type", "audio/mpeg")

	// _, err = io.Copy(w, audioFile)
	// if err != nil {
	// 	http.Error(w, "Unable to stream MP3", http.StatusInternalServerError)
	// }
	audioPath := "/home/iqthuc/Downloads/Nang-Tho-Hoang-Dung.mp3"
	// Lấy đường dẫn của tệp tin mà client muốn truy cập từ URL.
	http.ServeFile(w, r, audioPath)
}
func testApiMp3(w http.ResponseWriter, r *http.Request) {
	audioFile, err := os.ReadFile("/home/iqthuc/Downloads/Nang-Tho-Hoang-Dung.mp3")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Write(audioFile)
}

func main() {
	mainRouter := mux.NewRouter()

	homeRouter := mainRouter.PathPrefix("/home").Subrouter()

	homeRouter.HandleFunc("", getSomeData)
	homeRouter.HandleFunc("/test-mp3", testApiMp3)
	homeRouter.HandleFunc("/test-stream-api", testStreamMp3)

	mainServer := http.Server{
		Addr:    ":3060",
		Handler: mainRouter,
	}
	mainServer.ListenAndServe()
}
