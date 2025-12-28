package main

import (
	singleinstance "Everydo/internal/single_instance"
	"Everydo/internal/tray"
	"context"
	"embed"
	"log"
	"log/slog"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

var (
	doCloseApp bool = false
)

func main() {
	app := NewApp()

	ok := singleinstance.CheckSingleInstance(func() {
		runtime.WindowShow(app.ctx)
		runtime.Show(app.ctx)
	})

	if !ok {
		return
	}

	// Передаем контекст напрямую, без quitChan

	err := wails.Run(&options.App{
		Title:  "Everydo",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnDomReady: func(ctx context.Context) {
			go tray.SetupTray(ctx, &doCloseApp)
		},
		OnStartup: app.startup,
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			slog.Info("OnBeforeClose")
			runtime.WindowHide(ctx)
			if !doCloseApp {
				return true
			}
			return false
		},
		OnShutdown: func(ctx context.Context) {
			// Явно завершаем трей при закрытии приложения
			tray.Quit()
			log.Println("Application shutdown complete")
		},
		Bind: []interface{}{app},
	})

	if err != nil {
		panic(err)
	}
}
