package router

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

// Config -
type Config struct {
	Readme         `json:"readme"`
	VuepressConfig `json:"config"`
}

// Readme -
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

// VuepressConfig -
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
		Sidebar []struct {
			RouteName string `json:"routeName"`
			Items     []struct {
				Title       string   `json:"title"`
				Collapsable bool     `json:"collapsable"`
				Children    []string `json:"children"`
			} `json:"items"`
		} `json:"sidebar"`
	} `json:"themeConfig"`
}

// InitVuepress -
func InitVuepress(c *gin.Context) {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config Config
	json.Unmarshal(byteValue, &config)

	CreateReadme(c, config.Readme)

	runCommand(c, "./tool/test.sh")
}

// CreateReadme -
func CreateReadme(c *gin.Context, config Readme) {
	tmpl, err := template.ParseFiles("./README.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	f, err := os.Create("./README.md")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = tmpl.Execute(f, config)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}

func runCommand(c *gin.Context, execPath string) {
	args := make([]string, 0)
	args = append(args, "-c")
	args = append(args, execPath)
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", args...)

	//显示运行的命令
	c.String(200, "%s", cmd.Args)
	//

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		c.String(200, "Error starting command: %s......", err.Error())
		return
	}

	go asyncLog(c, stdout)
	go asyncLog(c, stderr)

	if err := cmd.Wait(); err != nil {
		c.String(200, "Error waiting for command execution: %s......", err.Error())
		return
	}
	return
}

func asyncLog(c *gin.Context, reader io.ReadCloser) error {
	cache := "" //缓存不足一行的日志信息
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if num > 0 {
			b := buf[:num]
			s := strings.Split(string(b), "\n")
			line := strings.Join(s[:len(s)-1], "\n") //取出整行的日志
			c.String(200, "%s%s\n", cache, line)
			cache = s[len(s)-1]
		}
	}
}
