package handleastrology

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func secondPage(w fyne.Window) {
	radioL := widget.NewLabel("Choose your status")
	radioW := widget.NewRadioGroup(
		[]string{"Single", "Relationship", "Married"},
		radioFunc,
	)

	radioContainer := container.New(layout.NewCenterLayout(), radioW)

	w.SetContent(container.NewVBox(
		radioL,
		radioContainer,
		layout.NewSpacer(),
		widget.NewButton("Back", back(w)),
	),
	)
}

func back(w fyne.Window) func() {
	return func() {
		firstPage(w)
	}
}

func radioFunc(s string) {
	player.married = false
	player.relationship = false
	player.single = false

	if s == "Single" {
		player.single = true
		return
	}

	if s == "Relationship" {
		player.married = true
		return
	}

	if s == "Married" {
		player.married = true
	}

}
