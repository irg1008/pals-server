package templates

import (
	"bytes"
	"html/template"
)

func parseTemplate[T any](file string, data T) (string, error) {
	t, err := template.ParseFiles(file)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

type ConfirmUserData struct {
	Name string
	URL  string
}

func GetConfirm(name string, url string) (string, error) {
	data := ConfirmUserData{name, url}
	return parseTemplate("templates/confirm.html", data)
}

func GetReset(name string, url string) (string, error) {
	data := ConfirmUserData{name, url}
	return parseTemplate("templates/reset.html", data)
}
