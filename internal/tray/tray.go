package tray

import (
	"context"
	"log"
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
func SetupTray(ctx *context.Context) {
	go func() {
		// Гарантируем однократный запуск трей
		once.Do(func() {
			systray.Run(onReady(ctx), onExit)
		})
	}()
}

// onReady настраивает элементы трея
func onReady(ctx *context.Context) func() {
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
					wruntime.WindowShow(*ctx)
				case <-mQuit.ClickedCh:
					// Корректно завершаем приложение
					wruntime.Quit(*ctx)
					return
				}
			}
		}()
	}
}

func onExit() {
	log.Println("Tray exited")
}

// Quit принудительно завершает трей
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
