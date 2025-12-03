package astrology

import (
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
)

var (
	resultLoadingSpeed = 20 * time.Millisecond
)

func resultPage() func() {
	return func() {
		descL := widget.NewLabel("Guessing your personality...")
		loading := widget.NewProgressBar()

		config.Cfg.Window.SetContent(container.NewVBox(
			descL,
			layout.NewSpacer(),
			loading,
		))
		config.Cfg.Window.Show()

		//StartLoading
		go func() {
			for i := 0.0; i <= 1.0; i += 0.01 {
				fyne.Do(func() { loading.SetValue(i) })
				time.Sleep(resultLoadingSpeed)
			}

			//Show results
			descL = widget.NewLabel("Your personality:")
			result := getResult()

			fyne.Do(func() {
				config.Cfg.Window.SetContent(container.NewVBox(
					descL,
					result,
					layout.NewSpacer(),
					widget.NewSeparator(),
					widget.NewButton("See interesting information", func() { rickrollOrRestore() }),
					config.HomeExitButtons,
				))
			})
		}()
	}
}

func getResult() *widget.Label {
	zodiac := strings.ToLower(config.User.ZodiacSign.Text)
	result := config.Roasts[zodiac][config.User.Status]

	resultL := widget.NewLabel(result)
	resultL.Alignment = fyne.TextAlignCenter
	return resultL
}
