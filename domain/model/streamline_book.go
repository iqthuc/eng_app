package model

type StreamlineBook struct {
	Id           string `json:"id"`
	OrdinalId    int    `json:"ordinal_id"`
	Title        string `json:"title"`
	CollectionId string `json:"collection_id"`
}

type StreamlineLesson struct {
	Id        int        `json:"id"`
	BookId    string     `json:"book_id"`
	GrammarId NullInt64  `json:"grammer_id"`
	OrdinalId int        `json:"ordinal_id"`
	Title     string     `json:"title"`
	Audio     NullString `json:"audio"`
	Html      string     `json:"html"`
	Vocab     string     `json:"vocab"`
}

type StreamlineLessonWithTitleBook struct {
	StreamlineLesson
	TitleBook string `json:"title_book"`
}
