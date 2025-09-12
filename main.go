package main

import (
	"os"

	"github.com/ifrunruhin12/inventory/data"
	"github.com/ifrunruhin12/inventory/templates"
	"github.com/ifrunruhin12/inventory/utils"
)

func main() {
	tmplMgr, err := templates.NewTemplateManager("templates", "report.tmpl")
	if err != nil {
		utils.Logger.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	report := data.DefaultReportData()

	err = tmplMgr.Execute(os.Stdout, "report.tmpl", report)
	if err != nil {
		utils.Logger.Error("Failed to execute template", "error", err)
		os.Exit(1)
	}
}
