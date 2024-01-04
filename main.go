package main

import (
	"fmt"
	"har-converter/harlogconverter"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	errs := ""
	directory := "/Users/dengzhehang/Code/test/har/iac"
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			harJson, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
				return err
			}
			c := harlogconverter.HarConverter{}
			c.Parse(harJson)
			errReq := c.FindErrReq()
			if errReq != "" {
				errs += errReq
			}
			c.GenIdeaHttpRequest(true)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	d := "/Users/dengzhehang/Code/test/har/iac.log"
	err2 := os.WriteFile(d, []byte(errs), 0644)
	if err2 != nil {
		return
	}
}
