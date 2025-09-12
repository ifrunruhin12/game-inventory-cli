// Package templates handles template parsing, function registration and template execution.
package templates

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/ifrunruhin12/inventory/utils"
)

type TemplateManager struct {
	tmpl *template.Template
}

func NewTemplateManager(templateDir string, templateFile string) (*TemplateManager, error) {
	funcs := template.FuncMap{
		"pluralize":  utils.Pluralize,
		"capitalize": utils.Capitalize,
	}

	tmplPath := filepath.Join(templateDir, templateFile)
	tmpl, err := template.New("").Funcs(funcs).ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}

	return &TemplateManager{tmpl: tmpl}, nil
}

func (tm *TemplateManager) Execute(output *os.File, templateName string, data any) error {
	return tm.tmpl.ExecuteTemplate(output, templateName, data)
}
