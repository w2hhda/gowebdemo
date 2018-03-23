package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"encoding/json"
	"strings"
	"unicode"
	"math/rand"
)

type News struct {
	Title   string
	Content string
	Date    string
	Source  string
	Heat    int
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["json"] = scrape()
	c.ServeJSON()
}

func scrape() string {
	doc, err := goquery.NewDocument("http://www.jinse.com/lives")
	if err != nil {
		log.Fatal(err)
	}

	var newsArr []News
	doc.Find(".lost > li").Each(func(i int, li *goquery.Selection) {
		liveTime := strings.FieldsFunc(li.Find(".live-date").Text(), unicode.IsSpace)[0]
		liveInfo := strings.FieldsFunc(li.Find(".live-info").Text(), unicode.IsSpace)[0]
		liveTitle := strings.FieldsFunc(liveInfo[0:strings.Index(liveInfo, "】") + len("】")], unicode.IsSpace)[0]
		item := News{
			liveTitle,
			liveInfo,
			liveTime,
			"ICO 监管",
			rand.Intn(10000),
		}
		newsArr = append(newsArr, item)
	})

	result, err := json.Marshal(newsArr)
	if err != nil {
		fmt.Println("fail")
		return ""
	}

	return string(result)
}
