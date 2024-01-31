package controller

import (
	"eng_app_module/database"
	"eng_app_module/domain/model"
	"eng_app_module/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func GetLessons(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	limitParam := query.Get("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 10
	}
	db := database.GetDB()
	querySQL := fmt.Sprintf(`SELECT books.title , lessons.*
	FROM books
	join lessons on books.id = lessons.book_id
	LIMIT %d;`, limit)
	rows, err := db.Query(querySQL)
	if err != nil {
		utils.ReponseCommonError(response)
		log.Println(err)
		return
	}
	defer rows.Close()

	var lessonList []model.StreamlineLessonWithTitleBook

	for rows.Next() {
		var lesson model.StreamlineLessonWithTitleBook

		err := rows.Scan(
			&lesson.TitleBook,
			&lesson.Id,
			&lesson.BookId,
			&lesson.GrammarId,
			&lesson.OrdinalId,
			&lesson.Title,
			&lesson.Audio,
			&lesson.Html,
			&lesson.Vocab,
		)

		if err != nil {
			utils.ReponseCommonError(response)
			log.Println(err)
			return
		}
		lessonList = append(lessonList, lesson)
	}
	result := model.BaseData[[]model.StreamlineLessonWithTitleBook]{
		Message: "success",
		Code:    http.StatusOK,
		Status:  1,
		Data:    lessonList,
	}

	utils.ReponseData(response, result.ToResponseData())
}
func GetLessonDetail(response http.ResponseWriter, request *http.Request) {
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
