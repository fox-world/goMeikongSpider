package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"labix.org/v2/mgo"
	"log"
	"models"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func MeikongSpider(url string, num int, modelCollection *mgo.Collection) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.index-module.noborder-module>div.w970>ul.post.small-post").Each(func(i int, s *goquery.Selection) {
		name := s.Find("li>a.nickname").Text()
		name = strings.Replace(name, ".", "", -1)
		click := s.Find("li").Eq(2).Find("span").Text()
		href, _ := s.Find("div.cover>a").Attr("href")
		num = num + 1
		fmt.Println("模特", num, ":", name)
		fmt.Println("点击量:", click)
		url = "http://www.moko.cc" + href
		fmt.Println("个人主页:", url)
		images := getModelInfo(url, name)
		clicknum, _ := strconv.Atoi(click)

		err = modelCollection.Insert(&models.Model{Name: name, Click: clicknum, Page: url, Address: images})

		if err != nil {
			panic(err)
		} else {
			fmt.Println("-------------insert data for model", name, "success-------------------")
		}
	})

	node := doc.Find("p.page>a").Last()
	_, exists := node.Attr("hidefocus")
	if exists {
		next, _ := node.Attr("href")
		next = "http://www.moko.cc" + next
		MeikongSpider(next, num, modelCollection)
	}

}

func getModelInfo(url string, name string) []string {

	//利用数组切片来存储图片，避免需要预先定义数组的大小
	images := make([]string, 0)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	dir := "模特" + string(filepath.Separator) + name
	os.MkdirAll(dir, os.ModeDir)

	doc.Find("div.article>div.pic.dBd").Each(func(i int, s *goquery.Selection) {
		content, _ := s.Find("p.picBox").Html()
		re := regexp.MustCompile("src2=\"(\\S+)\"")
		src := re.FindStringSubmatch(content)[1]
		if strings.Contains(src, "?") {
			src = src[:strings.LastIndex(src, "?")]
		}
		name := src[strings.LastIndex(src, "/")+1:]
		saveFileFromUrl(src, dir, name)
		images = append(images, src)
	})

	return images
}

func saveFileFromUrl(url string, dir string, name string) {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()
	file, err := os.Create(dir + string(filepath.Separator) + name)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println("create image", name, "success!")
}
