package structure

type LoginStruct struct {
	Code int       `json:"code"`
	Data LoginData `json:"data"`
}

type LoginData struct {
	LoginToken string     `json:"login_token"`
	UserCode   string     `json:"user_code"`
	ReaderInfo ReaderInfo `json:"reader_info"`
	PropInfo   PropInfo   `json:"prop_info"`
	IsSetYoung string     `json:"is_set_young"`
}

type PropInfo struct {
	RESTGiftHlb    int `json:"rest_gift_hlb"`
	RESTVipHlb     int `json:"rest_vip_hlb"`
	RESTHlb        int `json:"rest_hlb"`
	RESTYp         int `json:"rest_yp"`
	RESTRecommend  int `json:"rest_recommend"`
	RESTTotalBlade int `json:"rest_total_blade"`
	RESTMonthBlade int `json:"rest_month_blade"`
}

type ReaderInfo struct {
	ReaderID       string        `json:"reader_id"`
	Account        string        `json:"account"`
	IsBind         string        `json:"is_bind"`
	IsBindQq       int           `json:"is_bind_qq"`
	IsBindWeixin   int           `json:"is_bind_weixin"`
	PhoneNum       string        `json:"phone_num"`
	Email          string        `json:"email"`
	ReaderName     string        `json:"reader_name"`
	AvatarURL      string        `json:"avatar_url"`
	AvatarThumbURL string        `json:"avatar_thumb_url"`
	BaseStatus     string        `json:"base_status"`
	ExpLV          string        `json:"exp_lv"`
	Gender         string        `json:"gender"`
	VipLV          int           `json:"vip_lv"`
	IsAuthor       string        `json:"is_author"`
	BookAge        int           `json:"book_age"`
	TagPrefer      []interface{} `json:"tag_prefer"`
	UsedDecoration []interface{} `json:"used_decoration"`
	VipEndTime     string        `json:"vip_end_time"`
}
