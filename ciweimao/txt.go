package ciweimao

import (
	"fmt"
	"log"
	"os"

	"github.com/cheggaaa/pb"
	"github.com/mitchellh/go-homedir"
	"github.com/zsakvo/Cirno-go/config"
	"github.com/zsakvo/Cirno-go/structure"
)

var txtBar *pb.ProgressBar
var cacheNoPaidTxt bool

func DownloadText(bid string) {
	cacheNoPaidTxt = config.Config.Extra.CacheNoPaid
	txtDetail := GetDetail(bid)
	fmt.Println(txtDetail.BookName, "/", txtDetail.AuthorName)
	txtName := txtDetail.BookName
	txtChapters := GetCatalog(bid)
	var chapters []structure.ChapterList
	if !cacheNoPaidTxt {
		for _, c := range txtChapters {
			if c.AuthAccess == 1 {
				chapters = append(chapters, c)
			}
		}
	} else {
		chapters = txtChapters
	}
	txtTotalCount := len(chapters)
	txtBar = pb.StartNew(txtTotalCount)
	txtContainer := make(map[string]string)
	txtc := make(chan chapterStruct, 409600)
	errc := make(chan structure.ChapterList, 102400)
	coroutine := config.Config.Extra.Coroutines
	txtChaptersArr := splitArray(chapters, coroutine)
	for _, cs := range txtChaptersArr {
		go getChapterText(cs, txtc, errc)
	}
	for {
		select {
		case t := <-txtc:
			txtContainer[t.Cid] = t.Text
			txtBar.Increment()
		case e := <-errc:
			go getChapterText([]structure.ChapterList{e}, txtc, errc)
		}
		if len(txtContainer) == len(chapters) {
			close(txtc)
			close(errc)
			break
		}
	}
	txtBar.Finish()
	fmt.Println("writing out files…")
	writeText(txtName, txtContainer, chapters)
	fmt.Println("download success!")
	os.Exit(0)
}

func writeText(bookName string, txtContainer map[string]string, chapters []structure.ChapterList) {
	bookText := ""
	dir, _ := homedir.Dir()
	expandedDir, _ := homedir.Expand(dir)
	cleanDestBook(expandedDir+"/Cirno/download/", bookName, "txt")
	counter := 1
	for _, chapter := range chapters {
		bookText += fmt.Sprintln("第", counter, "章:", chapter.ChapterTitle)
		bookText += txtContainer[chapter.ChapterID]
		counter += 1
	}
	fileName := expandedDir + "/Cirno/download/" + bookName + ".txt"
	if isExist(fileName) {
		os.Remove(fileName)
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = file.WriteString(bookText)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
}

func getChapterText(chapters []structure.ChapterList, txt chan chapterStruct, errc chan structure.ChapterList) {
	for _, chapter := range chapters {
		text := ""
		chapterInfo, err := GetContent(chapter.ChapterID)
		if err != nil {
			errc <- chapter
		} else {
			//text += chapterInfo.ChapterTitle
			//text += "\n\n"
			text += chapterInfo.TxtContent
			text += chapterInfo.AuthorSay
			text += "\n\n\n"
			txtstr := chapterStruct{text, chapter.ChapterID}
			txt <- txtstr
		}
	}
}
