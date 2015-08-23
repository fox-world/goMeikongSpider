package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Page struct {
	PageNo      int
	PageSize    int
	Next        int
	Previous    int
	TotalRecord int
	TotalPage   int
}

type Model struct {
	Number  int
	Name    string
	Job     string
	Click   int
	Page    string
	Address []string
}

type Config struct {
	DBuri string
}

func (config *Config) ReadConfig() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/src/config.json")
	if err != nil {
		fmt.Println("Open config file error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		fmt.Println("Decode config file error:", err)
	}
}

func (page *Page) CalculatePageInfo() {

	var previous, next, totalPage = page.PageNo - 1, page.PageNo + 1, 0

	if page.TotalRecord%page.PageSize == 0 {
		totalPage = page.TotalRecord / page.PageSize
	} else {
		totalPage = page.TotalRecord/page.PageSize + 1
	}

	if page.PageNo == 1 {
		previous = 1
	}
	if page.PageNo == page.TotalPage {
		next = page.TotalPage
	}

	page.Previous = previous
	page.Next = next
	page.TotalPage = totalPage
}
