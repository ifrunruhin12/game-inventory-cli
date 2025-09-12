// Package handlers handles template parsing, function registration and template execution.
package handlers

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/ifrunruhin12/inventory/utils"
)

type TemplateManager struct {
	tmpl *template.Template
}

func NewTemplateManager(templateDir string, templateFiles ...string) (*TemplateManager, error) {
	funcs := template.FuncMap{
		"pluralize":  utils.Pluralize,
		"capitalize": utils.Capitalize,
	}

	tmpl := template.New("").Funcs(funcs)

	if len(templateFiles) == 0 {
		pattern := filepath.Join(templateDir, "*.tmpl")
		var err error
		tmpl, err = tmpl.ParseGlob(pattern)
		if err != nil {
			return nil, err
		}
	} else {
		var templatePaths []string
		for _, file := range templateFiles {
			templatePaths = append(templatePaths, filepath.Join(templateDir, file))
		}
		var err error
		tmpl, err = tmpl.ParseFiles(templatePaths...)
		if err != nil {
			return nil, err
		}
	}

	return &TemplateManager{tmpl: tmpl}, nil
}

func (tm *TemplateManager) Execute(output *os.File, templateName string, data any) error {
	return tm.tmpl.ExecuteTemplate(output, templateName, data)
}
