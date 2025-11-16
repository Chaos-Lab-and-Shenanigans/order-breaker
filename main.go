package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	tappedfunctions "github.com/Chaos-Lab-and-Shenanigans/order-breaker/tapped_functions"
)

func main() {
	app := app.New()
	window := app.NewWindow("Hello")
	win2 := app.NewWindow("2nd window")

	hello := widget.NewLabel("Hello Fyne!")
	window.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	Fuck := widget.NewLabel("Awaiting response...")
	win2.SetContent(container.NewVBox(
		Fuck,
		widget.NewButton("Fuck jai?", tappedfunctions.TapFuckJai(Fuck)),
		&widget.Check{},
	))

	//window.ShowAndRun()
	win2.ShowAndRun()
}
