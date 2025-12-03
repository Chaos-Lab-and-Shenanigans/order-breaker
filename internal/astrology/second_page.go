package astrology

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
)

func secondPage() {
	radioL := widget.NewLabel("Choose your status")
	radioW := widget.NewRadioGroup(
		[]string{"Single", "Relationship", "Married"},
		radioFunc,
	)

	radioContainer := container.New(layout.NewCenterLayout(), radioW)

	nextB := widget.NewButton(
		"Next",
		func() {
			if check() {
				thirdPage()
			}
		},
	)

	backB := widget.NewButton(
		"Back",
		func() {
			firstPage()
		},
	)

	navigation := container.New(layout.NewGridLayout(2), backB, nextB)

	config.Cfg.Window.SetContent(container.NewVBox(
		radioL,
		radioContainer,
		layout.NewSpacer(),
		navigation,
	),
	)
}

func check() bool {
	if config.User.ZodiacSign == nil {
		return false
	}

	switch config.User.Status {
	case config.IsSingle:
		return true
	case config.InRelationship:
		return true
	case config.IsMarried:
		return true
	}

	return false
}

func radioFunc(s string) {
	if s == "Single" {
		config.User.Status = config.IsSingle
		return
	}

	if s == "Relationship" {
		config.User.Status = config.IsMarried
		return
	}

	if s == "Married" {
		config.User.Status = config.IsMarried
	}

}
