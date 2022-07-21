package structure

type SearchStruct struct {
	Code int        `json:"code"`
	Data SearchData `json:"data"`
}

type SearchData struct {
	TagList     []SearchTagList `json:"tag_list"`
	BookList    []BookList      `json:"book_list"`
	BooklistNum int             `json:"booklist_num"`
}

type BookList struct {
	BookID          string `json:"book_id"`
	BookName        string `json:"book_name"`
	Description     string `json:"description"`
	BookSrc         string `json:"book_src"`
	CategoryName    string `json:"category_name"`
	Tag             string `json:"tag"`
	TotalWordCount  string `json:"total_word_count"`
	UpStatus        string `json:"up_status"`
	IsPaid          string `json:"is_paid"`
	Discount        int    `json:"discount"`
	DiscountEndTime string `json:"discount_end_time"`
	Cover           string `json:"cover"`
	AuthorName      string `json:"author_name"`
	Uptime          string `json:"uptime"`
	Newtime         string `json:"newtime"`
	ReviewAmount    string `json:"review_amount"`
	ChapterAmount   string `json:"chapter_amount"`
	TotalTsukkomi   string `json:"total_tsukkomi"`
	IsOriginal      int    `json:"is_original"`
	RewardAmount    string `json:"reward_amount"`
	TotalClick      int    `json:"total_click"`
	MonthClick      int    `json:"month_click"`
	WeekClick       int    `json:"week_click"`
	TotalRecommend  string `json:"total_recommend"`
	MonthRecommend  string `json:"month_recommend"`
	WeekRecommend   string `json:"week_recommend"`
	TotalFavor      string `json:"total_favor"`
	MonthFavor      string `json:"month_favor"`
	WeekFavor       string `json:"week_favor"`
	CurrentYp       string `json:"current_yp"`
	TotalYp         string `json:"total_yp"`
	CurrentBlade    string `json:"current_blade"`
	TotalBlade      string `json:"total_blade"`
	WeekFansValue   string `json:"week_fans_value"`
	MonthFansValue  string `json:"month_fans_value"`
	TotalFansValue  string `json:"total_fans_value"`
	LastChapterInfo struct {
	} `json:"last_chapter_info"`
	SignStatus string `json:"sign_status"`
	ShortNovel string `json:"short_novel"`
}

type SearchLastChapterInfo struct {
	ChapterID         string `json:"chapter_id"`
	BookID            string `json:"book_id"`
	ChapterIndex      string `json:"chapter_index"`
	ChapterTitle      string `json:"chapter_title"`
	Uptime            string `json:"uptime"`
	Mtime             string `json:"mtime"`
	RecommendBookInfo string `json:"recommend_book_info"`
}

type SearchTagList struct {
	TagName string `json:"tag_name"`
	Num     string `json:"num"`
}
