package structure

type DetailStruct struct {
	Code int        `json:"code"`
	Data DetailData `json:"data"`
}

type DetailData struct {
	BookInfo      BookInfo        `json:"book_info"`
	IsInshelf     int             `json:"is_inshelf"`
	IsBuy         int             `json:"is_buy"`
	UpReaderInfo  UpReaderInfo    `json:"up_reader_info"`
	RelatedList   []BookInfo      `json:"related_list"`
	HotReviewList []HotReviewList `json:"book_shortage_reommend_list"`
}

type BookInfo struct {
	BookID          string          `json:"book_id"`
	BookName        string          `json:"book_name"`
	Description     string          `json:"description"`
	BookSrc         string          `json:"book_src"`
	CategoryName    string          `json:"category_name"`
	Tag             string          `json:"tag"`
	TotalWordCount  string          `json:"total_word_count"`
	UpStatus        string          `json:"up_status"`
	IsPaid          string          `json:"is_paid"`
	Discount        int             `json:"discount"`
	DiscountEndTime string          `json:"discount_end_time"`
	Cover           string          `json:"cover"`
	AuthorName      string          `json:"author_name"`
	Uptime          string          `json:"uptime"`
	Newtime         string          `json:"newtime"`
	ReviewAmount    string          `json:"review_amount"`
	ChapterAmount   string          `json:"chapter_amount"`
	TotalTsukkomi   string          `json:"total_tsukkomi"`
	IsOriginal      int             `json:"is_original"`
	RewardAmount    string          `json:"reward_amount"`
	TotalClick      int             `json:"total_click"`
	MonthClick      int             `json:"month_click"`
	WeekClick       int             `json:"week_click"`
	TotalRecommend  string          `json:"total_recommend"`
	MonthRecommend  string          `json:"month_recommend"`
	WeekRecommend   string          `json:"week_recommend"`
	TotalFavor      string          `json:"total_favor"`
	MonthFavor      string          `json:"month_favor"`
	WeekFavor       string          `json:"week_favor"`
	CurrentYp       string          `json:"current_yp"`
	TotalYp         string          `json:"total_yp"`
	CurrentBlade    string          `json:"current_blade"`
	TotalBlade      string          `json:"total_blade"`
	WeekFansValue   string          `json:"week_fans_value"`
	MonthFansValue  string          `json:"month_fans_value"`
	TotalFansValue  string          `json:"total_fans_value"`
	LastChapterInfo LastChapterInfo `json:"last_chapter_info"`
	SignStatus      string          `json:"sign_status"`
	ShortNovel      string          `json:"short_novel"`
}

type LastChapterInfo struct {
	ChapterId    string `json:"chapter_id"`
	BookId       string `json:"book_id"`
	ChapterIndex string `json:"chapter_index"`
	ChapterTitle string `json:"chapter_title"`
	Uptime       string `json:"uptime"`
}

type HotReviewList struct {
	ReviewId      string `json:"review_id"`
	BookId        string `json:"book_id"`
	BookName      string `json:"book_name"`
	Type          string `json:"type"`
	Title         string `json:"title"`
	ReviewContent string `json:"review_content"`
	Rank          string `json:"rank"`
	ReaderInfo    struct {
		ReaderId       string        `json:"reader_id"`
		ReaderName     string        `json:"reader_name"`
		AvatarUrl      string        `json:"avatar_url"`
		AvatarThumbUrl string        `json:"avatar_thumb_url"`
		BaseStatus     string        `json:"base_status"`
		ExpLv          string        `json:"exp_lv"`
		Gender         string        `json:"gender"`
		VipLv          int           `json:"vip_lv"`
		IsAuthor       string        `json:"is_author"`
		IsFollowing    string        `json:"is_following"`
		UsedDecoration []interface{} `json:"used_decoration"`
		IsInBlacklist  string        `json:"is_in_blacklist"`
		BookFansValue  string        `json:"book_fans_value"`
	} `json:"reader_info"`
	LikeAmount    string `json:"like_amount"`
	CommentAmount string `json:"comment_amount"`
	IsLike        int    `json:"is_like"`
	Ctime         string `json:"ctime"`
}

type UpReaderInfo struct {
	ReaderID       string        `json:"reader_id"`
	Account        string        `json:"account"`
	ReaderName     string        `json:"reader_name"`
	AvatarURL      string        `json:"avatar_url"`
	AvatarThumbURL string        `json:"avatar_thumb_url"`
	BaseStatus     string        `json:"base_status"`
	ExpLv          string        `json:"exp_lv"`
	Gender         string        `json:"gender"`
	VipLv          int           `json:"vip_lv"`
	IsAuthor       string        `json:"is_author"`
	IsFollowing    string        `json:"is_following"`
	UsedDecoration []interface{} `json:"used_decoration"`
	IsInBlacklist  string        `json:"is_in_blacklist"`
}
