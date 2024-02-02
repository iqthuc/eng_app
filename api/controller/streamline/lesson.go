package controller

import (
	"eng_app_module/database"
	"eng_app_module/domain/model"
	"eng_app_module/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetLessons(response http.ResponseWriter, request *http.Request) {
	db := database.GetDB()
	querySQL := "SELECT books.title , lessons.* FROM books join lessons on books.id = lessons.book_id"

	query := request.URL.Query()

	bookIdParam := query.Get("book_id")
	if bookIdParam != "" {
		querySQL += fmt.Sprintf(" WHERE books.id = '%s'", bookIdParam)
	}

	limitParam, err := strconv.Atoi(query.Get("limit"))
	if limitParam != 0 && err == nil {
		querySQL += fmt.Sprintf(" limit %d", limitParam)
	}

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
	querySQL := "SELECT books.title , lessons.* FROM books join lessons on books.id = lessons.book_id"

	// xử lý lấy id trong path
	path := request.URL.Path
	parts := strings.Split(path, "/")
	var id string
	for i, part := range parts {
		if part == "lessons" && i < (len(parts)-1) {
			id = string(parts[i+1])
			break
		}
	}

	if id != "" {
		querySQL += fmt.Sprintf(" WHERE lessons.id = '%s'", id)
	}

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
