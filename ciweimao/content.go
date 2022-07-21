package ciweimao

import (
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
	"github.com/zsakvo/Cirno-go/structure"
	"github.com/zsakvo/Cirno-go/util"
)

func GetContent(cid string) (structure.ChapterInfo, error) {
	var err error
	var res []byte
	paras := req.Param{
		"chapter_id": cid,
	}
	res, err = util.Get("/chapter/get_chapter_info", paras)
	if err != nil {
		return structure.ChapterInfo{}, err
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var result structure.ContentStruct
	err = json.Unmarshal(res, &result)
	if err != nil {
		return structure.ChapterInfo{}, err
	}
	return result.Data.ChapterInfo, nil
}
