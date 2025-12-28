package tray

import (
	"context"
	"log"
	"runtime"

	_ "embed"

	"github.com/getlantern/systray"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed icon.ico
var iconData []byte

//go:embed linux_icon.png
var linuxIconData []byte

func SetupTray(ctx *context.Context) {
	systray.Run(func() {
		systray.SetIcon(getIcon())
		systray.SetTitle("Everydo")
		systray.SetTooltip("Everydo | Menu")

		mShow := systray.AddMenuItem("Открыть", "Показать окно")
		mQuit := systray.AddMenuItem("Выход", "Закрыть приложение")

		go func() {
			for {
				select {
				case <-mShow.ClickedCh:
					openWindow(*ctx)
				case <-mQuit.ClickedCh:
					systray.Quit()
					wruntime.Quit(*ctx)
					return
				}
			}
		}()
	}, func() {
		log.Println("Tray exited")
	})
}

func getIcon() []byte {
	switch runtime.GOOS {
	case "windows":
		return iconData
	case "linux":
		return linuxIconData
	case "darwin":
		return linuxIconData
	default:
		return nil
	}
}

func openWindow(ctx context.Context) {
	wruntime.Show(ctx)
	wruntime.WindowShow(ctx)
}
