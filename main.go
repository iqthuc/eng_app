package main

import (

	// "log"
	database "eng_app_module/config"
	"fmt"
	"net/http"

	// "os"

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

func main() {
	mainRouter := mux.NewRouter()

	homeRouter := mainRouter.PathPrefix("/home").Subrouter()
	imageRouter := mainRouter.PathPrefix("/photos").Subrouter()

	audioDirectory := "/home/iqthuc/Downloads"

	homeRouter.PathPrefix("/audio").Handler(http.StripPrefix("/home/audio", http.FileServer(http.Dir(audioDirectory))))

	imageRouter.PathPrefix("").Handler(http.StripPrefix("/photos", http.FileServer(http.Dir(audioDirectory))))

	mainServer := http.Server{
		Addr:    ":3060",
		Handler: mainRouter,
	}
	mainServer.ListenAndServe()
}
