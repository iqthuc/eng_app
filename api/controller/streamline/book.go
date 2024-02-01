package controller

import (
	"eng_app_module/database"
	"eng_app_module/domain/model"
	"eng_app_module/internal/utils"
	"log"
	"net/http"
)

func GetBooks(response http.ResponseWriter, request *http.Request) {
	db := database.GetDB()
	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		utils.ReponseCommonError(response)
		log.Println(err)
		return
	}
	defer rows.Close()

	var bookList []model.StreamlineBook

	for rows.Next() {
		var book model.StreamlineBook
		err := rows.Scan(&book.Id, &book.OrdinalId, &book.Title, &book.CollectionId)
		if err != nil {
			utils.ReponseCommonError(response)
			log.Println(err)
			return
		}
		bookList = append(bookList, book)
	}
	result := model.BaseData[[]model.StreamlineBook]{
		Message: "success",
		Code:    http.StatusOK,
		Status:  1,
		Data:    bookList,
	}

	utils.ReponseData(response, result.ToResponseData())
}
