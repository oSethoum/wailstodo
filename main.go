package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "wailstodo",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 26, G: 27, B: 30},
		OnStartup:        app.startup,
		MinWidth:         420,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
