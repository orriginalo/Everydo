package main

import (
	singleinstance "Everydo/internal/single_instance"
	"Everydo/internal/tray"
	"context"
	"embed"
	"log"
	"log/slog"
	"runtime"

	"github.com/gen2brain/beeep"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

var (
	doCloseApp bool = false
)

func main() {
	beeep.AppName = "Everydo"

	app := NewApp()

	ok := singleinstance.CheckSingleInstance(func() {
		wruntime.WindowShow(app.ctx)
		wruntime.Show(app.ctx)
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
			if runtime.GOOS != "darwin" {
				go tray.SetupTray(ctx, &doCloseApp)
			}
		},
		OnStartup: app.startup,
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			slog.Info("OnBeforeClose")
			wruntime.WindowHide(ctx)
			return !doCloseApp
		},
		OnShutdown: func(ctx context.Context) {
			// Явно завершаем трей при закрытии приложения
			if runtime.GOOS == "darwin" {
				tray.Quit()
			}
			log.Println("Application shutdown complete")
		},
		Bind: []interface{}{app},
	})

	if err != nil {
		panic(err)
	}
}
