package url_handle

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func renderTemplate(qrType, qrTemplate string) (*template.Template, error) {

	layoutFiles, err := filepath.Glob(view + "/layout/*html")

	if err != nil {
		return nil, err
	}
	index := template.New("layout")

	for _, f := range layoutFiles {
		index, err = index.ParseFiles(f)
		if err != nil {
			return nil, err
		}
	}

	var dir = strings.ToLower(qrType)
	var templateName = "default"
	folderTemplate, _ := ioutil.ReadDir(view + "/" + dir + "/")
	for _, f := range folderTemplate {
		if f.IsDir() {
			if f.Name() == qrTemplate {
				templateName = qrTemplate
			}
		}
	}

	cssFiles, err := filepath.Glob(view + "/" + dir + "/" + templateName + "/*css")

	if err != nil {
		return nil, err
	}

	var css = "{{define \"css\"}}"

	for _, f := range cssFiles {
		_, fn := filepath.Split(f)
		css += "<link rel=\"stylesheet\" href=\"/static/view/" + dir + "/" + templateName + "/" + fn + "\">"
	}

	css += "{{end}}"

	tempFiles, err := filepath.Glob(view + "/" + dir + "/" + templateName + "/*html")

	if err != nil {
		return nil, err
	}

	for _, f := range tempFiles {
		index, err = index.ParseFiles(f)
		if err != nil {
			return nil, err
		}
	}

	index, err = index.Parse(css)

	if err != nil {
		return nil, err
	}

	temp := template.Must(index, err)

	return temp, nil
}
