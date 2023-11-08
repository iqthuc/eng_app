package main

import (
	"eng_app_module/api/route"
	database "eng_app_module/database"
	"fmt"
	"net/http"
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
	// mainRouter := mux.NewRouter()

	// homeRouter := mainRouter.PathPrefix("/home").Subrouter()
	// imageRouter := mainRouter.PathPrefix("/media").Subrouter()

	// mediaDirectory := "/Users/qsoft028/Documents/english_database/AmericanStreamline/Departures"

	// homeRouter.PathPrefix("").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Welcome home")
	// })

	// imageRouter.PathPrefix("").Handler(http.StripPrefix("/media", http.FileServer(http.Dir(mediaDirectory))))

	// mainServer := http.Server{
	// 	Addr:    ":3060",
	// 	Handler: mainRouter,
	// }
	// mainServer.ListenAndServe()

	route.Setup()
}
