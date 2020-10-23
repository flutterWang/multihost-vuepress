package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

type Config struct {
	Readme         `json:"readme"`
	VuepressConfig `json:"config"`
}

type Readme struct {
	Home       bool   `json:"home"`
	HeroImage  string `json:"heroImage"`
	ActionText string `json:"actionText"`
	ActionLink string `json:"actionLink"`
	HeroText   string `json:"heroText"`
	Sidebar    string `json:"sidebar"`
	Features   []struct {
		Title   string `json:"title"`
		Details string `json:"details"`
	} `json:"features"`
	Footer string `json:"footer"`
}

type VuepressConfig struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Head        []struct {
		TagName string `json:"tagName"`
		Attrs   []struct {
			AttrNameK  string `json:"attrNameK"`
			AttrNameV  string `json:"attrNameV"`
			AttrValueK string `json:"attrValueK"`
			AttrValueV string `json:"attrValueV"`
		} `json:"attrs"`
	} `json:"head"`
	Markdown struct {
		LineNumbers bool `json:"lineNumbers"`
	} `json:"markdown"`
	ThemeConfig struct {
		Nav []struct {
			Text string `json:"text"`
			Link string `json:"link"`
		} `json:"nav"`
	} `json:"themeConfig"`
}

func InitVuepress(c *gin.Context) {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config Config
	json.Unmarshal(byteValue, &config)

	CreateReadme(config.Readme)
	CreateConfig(config.VuepressConfig)

}

func CreateReadme(config Readme) {
	fileName := "README.md"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	file.WriteString("--- \n")

	home := fmt.Sprintf("home: %v \n", config.Home)
	file.WriteString(home)

	heroImage := fmt.Sprintf("heroImage: %s \n", config.HeroImage)
	file.WriteString(heroImage)

	actionText := fmt.Sprintf("actionText: %s \n", config.ActionText)
	file.WriteString(actionText)

	actionLink := fmt.Sprintf("actionLink: %s \n", config.ActionLink)
	file.WriteString(actionLink)

	heroText := fmt.Sprintf("heroText: %s \n", config.HeroText)
	file.WriteString(heroText)

	sidebar := fmt.Sprintf("sidebar: %s \n", config.Sidebar)
	file.WriteString(sidebar)

	features := fmt.Sprintf("features: \n")
	file.WriteString(features)
	for _, feature := range config.Features {
		title := fmt.Sprintf("- title: %s \n", feature.Title)
		file.WriteString(title)
		details := fmt.Sprintf("  details: %s \n", feature.Details)
		file.WriteString(details)
	}

	footer := fmt.Sprintf("footer: %s \n", config.Footer)
	file.WriteString(footer)

	file.WriteString("--- \n")
}

func CreateConfig(config VuepressConfig) {
	fileName := "config.js"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()

	file.WriteString("module.exports = {")

	title := fmt.Sprintf(" title: '%s',", config.Title)
	file.WriteString(title)

	description := fmt.Sprintf(" description: '%s',", config.Description)
	file.WriteString(description)

	headBefore := fmt.Sprintf(" head: [")
	file.WriteString(headBefore)
	for _, head := range config.Head {
		writeArrBefore(file)
		tagName := fmt.Sprintf("    '%s',", head.TagName)
		file.WriteString(tagName)

		writeObjectBefore(file)

		for _, attr := range head.Attrs {
			attrName := fmt.Sprintf("    %s: '%v',", attr.AttrNameK, attr.AttrNameV)
			file.WriteString(attrName)
			attrValue := fmt.Sprintf("    %s: '%v',", attr.AttrValueK, attr.AttrValueV)
			file.WriteString(attrValue)
		}

		writeObjectAfter(file)
		writeArrAfter(file)
	}
	writeArrAfter(file)

	markdown := fmt.Sprintf(" markdown:{lineNumbers: %v,},", config.Markdown.LineNumbers)
	file.WriteString(markdown)

	themeConfig := fmt.Sprintf(" themeConfig:{nav:")
	file.WriteString(themeConfig)
	writeArrBefore(file)
	for _, nav := range config.ThemeConfig.Nav {
		writeObjectBefore(file)

		text := fmt.Sprintf(" text: '%s',", nav.Text)
		file.WriteString(text)
		link := fmt.Sprintf(" link: '%s',", nav.Link)
		file.WriteString(link)

		writeObjectAfter(file)
	}
	writeArrAfter(file)
	writeObjectAfter(file)

	file.WriteString("}")
}

func writeArrBefore(file *os.File) {
	before := fmt.Sprintf("[")
	file.WriteString(before)
}

func writeArrAfter(file *os.File) {
	after := fmt.Sprintf("],")
	file.WriteString(after)
}

func writeObjectBefore(file *os.File) {
	before := fmt.Sprintf("{")
	file.WriteString(before)
}

func writeObjectAfter(file *os.File) {
	after := fmt.Sprintf("},")
	file.WriteString(after)
}
