//go:build !darwin

package tray

import (
	"context"
	"log"
	"log/slog"
	"runtime"
	"sync"

	_ "embed"

	"github.com/getlantern/systray"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed icon.ico
var iconData []byte

//go:embed linux_icon.png
var linuxIconData []byte

var (
	once     sync.Once
	quitOnce sync.Once
)

// SetupTray инициализирует системный трей
func SetupTray(ctx context.Context, doCloseApp *bool) {
	go func() {
		// Гарантируем однократный запуск трей
		once.Do(func() {
			systray.Run(onReady(ctx, doCloseApp), onExit)
		})
	}()
}

// onReady настраивает элементы трея
func onReady(ctx context.Context, doCloseApp *bool) func() {
	return func() {
		systray.SetIcon(getIcon())
		systray.SetTitle("Everydo")
		systray.SetTooltip("Everydo | Menu")

		mShow := systray.AddMenuItem("Открыть", "Показать окно")
		mQuit := systray.AddMenuItem("Выход", "Закрыть приложение")

		go func() {
			for {
				select {
				case <-mShow.ClickedCh:
					slog.Info("Show.Clicked")
					wruntime.WindowShow(ctx)
				case <-mQuit.ClickedCh:
					slog.Info("Quit.Clicked")
					*doCloseApp = true
					// Корректно завершаем приложение
					wruntime.Quit(ctx)
					return
				}
			}
		}()
	}
}

func onExit() {
	log.Println("Tray exited")
}

func Quit() {
	quitOnce.Do(func() {
		systray.Quit()
		log.Println("Tray force quit")
	})
}

func getIcon() []byte {
	switch runtime.GOOS {
	case "windows":
		return iconData
	case "linux", "darwin":
		return linuxIconData
	default:
		return linuxIconData
	}
}
