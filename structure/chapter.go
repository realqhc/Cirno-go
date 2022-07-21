package structure

type ChapterStruct struct {
	Code int         `json:"code"`
	Data ChapterData `json:"data"`
}

type ChapterData struct {
	ChapterList     []ChapterList `json:"chapter_list"`
	MaxUpdateTime   int           `json:"max_update_time"`
	MaxChapterIndex int           `json:"max_chapter_index"`
}

type ChapterList struct {
	ChapterID      string `json:"chapter_id"`
	BaseStatus     string `json:"base_status"`
	ChapterIndex   string `json:"chapter_index"`
	ChapterTitle   string `json:"chapter_title"`
	WordCount      string `json:"word_count"`
	TsukkomiAmount string `json:"tsukkomi_amount"`
	IsPaid         string `json:"is_paid"`
	Uptime         string `json:"uptime"`
	IsValid        string `json:"is_valid"`
	AuthAccess     int    `json:"auth_access"`
	DeleteStatus   int    `json:"delete_status"`
}
