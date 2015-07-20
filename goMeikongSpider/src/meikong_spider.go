package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	meikongSpider()
}

func meikongSpider() {
	doc, err := goquery.NewDocument("http://www.moko.cc/channels/post/23/1.html")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.index-module.noborder-module>div.w970>ul.post.small-post").Each(func(i int, s *goquery.Selection) {
		name := s.Find("li>a.nickname").Text()
		name = strings.Replace(name, ".", "", -1)
		click := s.Find("li").Eq(2).Find("span").Text()
		href, _ := s.Find("div.cover>a").Attr("href")
		fmt.Println("模特", i+1, ":", name)
		fmt.Println("点击量:", click)
		url := "http://www.moko.cc" + href
		fmt.Println("个人主页:", url)
		getModelInfo(url, name)
		fmt.Println("-----------------------")
	})

}

func getModelInfo(url string, name string) {

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
	})

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
