package sqltemplate

import (
	"fmt"

	"github.com/nikolalohinski/gonja/v2"
	"github.com/nikolalohinski/gonja/v2/exec"
)

type Template struct {
	template *exec.Template
}

func MustParse(name, source string) *Template {
	template, err := gonja.FromString(source)
	if err != nil {
		panic(fmt.Sprintf("parse SQL template %s: %v", name, err))
	}
	return &Template{template: template}
}

func (t *Template) Render(values map[string]any) (string, error) {
	query, err := t.template.ExecuteToString(exec.NewContext(values))
	if err != nil {
		return "", fmt.Errorf("render SQL template: %w", err)
	}
	return query, nil
}
