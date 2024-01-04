package harlogconverter

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const HEADERS = "\nAccept: application/json" +
	"\nAccess-Token: Bearer {{auth_token}}" +
	"\nAuthorization: {{auth_token}}" +
	"\nContent-Type: application/json" +
	"\n\n"

const ASSET = "> {%\n    client.test(\"status->success\", function() {\n        client.assert(response.status === 200, \"Response status is not 200\");\n        client.assert(response.body.status === 'success', \"Response status is not success\");\n    });\n%}"

type HarConverter struct {
	Har  Har
	Urls map[string]int
}

func (c *HarConverter) Parse(harDoc []byte) {
	err := json.Unmarshal(harDoc, &c.Har)
	if err != nil {
		fmt.Println("解析Har Doc出现错误：", err)
		os.Exit(1)
	}

	c.Urls = make(map[string]int)
	fmt.Println("解析出请求数量为：", len(c.Har.Log.Entries))
}

func (c *HarConverter) FindErrReq() string {
	c.Urls = make(map[string]int)
	var s string
	for _, entry := range c.Har.Log.Entries {
		method := strings.ToUpper(entry.Request.Method)
		url := entry.Request.Url
		url = strings.Trim(url, "?")
		payload := entry.Request.PostData.Text
		repBody := entry.Response.Content.Text

		if strings.HasSuffix(url, "cmp/core/api/v1/web_access/log") || strings.HasSuffix(url, ".js") || strings.HasSuffix(url, ".svg") || strings.HasSuffix(url, ".png") {
			continue
		}

		// .png
		if strings.HasSuffix(url, ".js") || strings.HasSuffix(url, ".svg") || strings.HasSuffix(url, ".png") {
			continue
		}

		if method == "GET" {
			if c.Urls[url] == 1 {
				continue
			}
			c.Urls[url] = 1
		}

		if entry.Response.Status > 300 || strings.HasPrefix(repBody, "{\"status\":\"fail\"") || strings.HasPrefix(repBody, "\"isSuccess\": false") {
			//if entry.Response.Status > 300 || strings.Contains(repBody, "\"status\":\"fail\"") || strings.Contains(repBody, "\"isSuccess\": false") {
			s += fmt.Sprintln("#################################################################")
			s += fmt.Sprintln("错误的请求：", method, url)
			if strings.HasSuffix(url, "resource/pool/execute") || payload != "" {
				s += fmt.Sprintln("payload：", payload)
			}
			s += fmt.Sprintln("返回：", repBody)
			s += fmt.Sprintln("#################################################################")
			s += fmt.Sprintln("\n\n")
		}
	}
	return s
}

func (c *HarConverter) GenIdeaHttpRequest(badOnly bool) {
	c.Urls = make(map[string]int)
	var request string
	for _, entry := range c.Har.Log.Entries {
		method := strings.ToUpper(entry.Request.Method)
		url := entry.Request.Url
		url = strings.Trim(url, "?")

		// .png
		if strings.HasSuffix(url, ".js") || strings.HasSuffix(url, ".svg") || strings.HasSuffix(url, ".png") {
			continue
		}

		if badOnly {
			var skip bool
			repBody := entry.Response.Content.Text
			if entry.Response.Status < 300 {
				skip = true
			} else if strings.HasPrefix(repBody, "{\"status\":\"success\"") {
				skip = true
			} else if strings.HasPrefix(repBody, "\"isSuccess\": true") {
				skip = true
			}

			if skip {
				continue
			}
		}

		if method == "GET" {
			if c.Urls[url] == 1 {
				continue
			}
			c.Urls[url] = 1
		}

		request += fmt.Sprintf("### \n%s %s %s", method, url, HEADERS)

		if method == "POST" || method == "PUT" || method == "DELETE" {
			if entry.Request.PostData.Text != "" {
				request += entry.Request.PostData.Text + "\n\n"
			}
		}

		request += ASSET + "\n\n"
	}

	d := "/Users/dengzhehang/Code/test/har/test.http"
	err := os.WriteFile(d, []byte(request), 0644)
	if err != nil {
		return
	}
}
