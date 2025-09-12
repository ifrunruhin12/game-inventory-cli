package main

import (
	"os"

	"github.com/ifrunruhin12/inventory/data"
	"github.com/ifrunruhin12/inventory/handlers"
	"github.com/ifrunruhin12/inventory/utils"
)

func main() {
	utils.InitLogger()

	tmplMgr, err := handlers.NewTemplateManager("templates")
	if err != nil {
		utils.Logger.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	err = data.StartInteractiveSession(tmplMgr)
	if err != nil {
		utils.Logger.Error("Interactive session failed", "error", err)
		os.Exit(1)
	}
}
