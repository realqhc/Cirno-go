package structure

type ContentStruct struct {
	Code int         `json:"code"`
	Data ContentData `json:"data"`
}

type ContentData struct {
	ChapterInfo ChapterInfo `json:"chapter_info"`
}

type ChapterInfo struct {
	ChapterID      string `json:"chapter_id"`
	BookID         string `json:"book_id"`
	DivisionID     string `json:"division_id"`
	UnitHlb        int    `json:"unit_hlb"`
	ChapterIndex   string `json:"chapter_index"`
	ChapterTitle   string `json:"chapter_title"`
	AuthorSay      string `json:"author_say"`
	WordCount      string `json:"word_count"`
	Discount       string `json:"discount"`
	IsPaid         string `json:"is_paid"`
	AuthAccess     int    `json:"auth_access"`
	BuyAmount      string `json:"buy_amount"`
	TsukkomiAmount string `json:"tsukkomi_amount"`
	ReviewAmount   string `json:"review_amount"`
	TotalHlb       string `json:"total_hlb"`
	Uptime         string `json:"uptime"`
	BaseStatus     string `json:"base_status"`
	DeleteStatus   int    `json:"delete_status"`
	TxtContent     string `json:"txt_content"`
}

type T struct {
	Code int `json:"code"`
	Data struct {
		ChapterInfo struct {
			ChapterId      string `json:"chapter_id"`
			BookId         string `json:"book_id"`
			DivisionId     string `json:"division_id"`
			UnitHlb        int    `json:"unit_hlb"`
			ChapterIndex   string `json:"chapter_index"`
			ChapterTitle   string `json:"chapter_title"`
			AuthorSay      string `json:"author_say"`
			WordCount      string `json:"word_count"`
			Discount       string `json:"discount"`
			IsPaid         string `json:"is_paid"`
			AuthAccess     int    `json:"auth_access"`
			BuyAmount      string `json:"buy_amount"`
			TsukkomiAmount string `json:"tsukkomi_amount"`
			ReviewAmount   string `json:"review_amount"`
			TotalHlb       string `json:"total_hlb"`
			Uptime         string `json:"uptime"`
			BaseStatus     string `json:"base_status"`
			DeleteStatus   int    `json:"delete_status"`
			TxtContent     string `json:"txt_content"`
		} `json:"chapter_info"`
	} `json:"data"`
}
