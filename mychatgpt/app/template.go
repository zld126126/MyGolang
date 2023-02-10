package app

import (
	"html/template"
	"io/ioutil"
	"strings"
)

// LoadTemplate 执行命令: go-assets-builder templates -o assets.go -p app
func LoadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		// 可以用.tmpl .html
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
