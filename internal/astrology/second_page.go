package astrology

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func secondPage() {
	makeStatusFalse()
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

	cfg.Window.SetContent(container.NewVBox(
		radioL,
		radioContainer,
		layout.NewSpacer(),
		navigation,
	),
	)
}

func check() bool {
	if player.isMarried || player.inRelationship || player.isSingle {
		return true
	}
	return false
}

func radioFunc(s string) {
	makeStatusFalse()

	if s == "Single" {
		player.isSingle = true
		return
	}

	if s == "Relationship" {
		player.isMarried = true
		return
	}

	if s == "Married" {
		player.isMarried = true
	}

}

func makeStatusFalse() {
	player.isMarried = false
	player.inRelationship = false
	player.isSingle = false
}
