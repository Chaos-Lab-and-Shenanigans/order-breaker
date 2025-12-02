package astrology

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	loadingSpeed = 20 * time.Millisecond
)

func StartCC() func() {
	return func() {
		thirdPage()
	}
}

func thirdPage() {
	descL := widget.NewLabel("Compatibility check")
	descContainer := container.New(layout.NewCenterLayout(), descL) //Center this shit

	p1 := widget.NewLabel("Enter your date of birth:")
	p2 := widget.NewLabel("Enter the date of brith of person you're interested in:")

	//Making labels left sided
	p1Container := container.New(layout.NewVBoxLayout(), p1)
	p2Container := container.New(layout.NewVBoxLayout(), p2)

	p1Dob := widget.NewDateEntry()
	p2Dob := widget.NewDateEntry()

	checkStatusB := widget.NewButton("Check status", calcStatus(p1Dob, p2Dob))

	mainContainer := container.NewVBox(
		descContainer,
		p1Container,
		p1Dob,
		p2Container,
		p2Dob,
		layout.NewSpacer(),
		checkStatusB,
	)

	cfg.Window.SetContent(mainContainer)
}

func calcStatus(p1Dob *widget.DateEntry, p2Dob *widget.DateEntry) func() {
	return func() {
		var compatible bool
		t1 := p1Dob.Date
		t2 := p2Dob.Date

		if t1 == nil || t2 == nil {
			return
		}

		player.dob = p1Dob.Text
		y1, m1, d1 := t1.Date()
		y2, m2, d2 := t2.Date()

		//random shit to determine compatibility, even I didn't bother understanding it
		if (y1-3 < y2 && y2 < y1+3) && (m1-2 < m2 && m2 < m1+2) && (d1-5 < d2 && d2 < d1+5) {
			compatible = true
		}

		showCompatibility(compatible)
	}
}

func showCompatibility(compatible bool) {
	descL := widget.NewLabel("Checking compatibility...")

	homeB := widget.NewButton("Home", func() { cfg.RestartCh <- "restart" })
	exitB := widget.NewButton("Exit", func() { fyne.CurrentApp().Quit() })

	navigationButtons := container.New(layout.NewGridLayout(2), exitB, homeB)
	loading := getLoading(compatible, navigationButtons)

	cfg.Window.SetContent(container.NewVBox(
		descL,
		layout.NewSpacer(),
		loading,
		widget.NewSeparator(),
		navigationButtons,
	))
}

func getLoading(compatible bool, navigationButtons *fyne.Container) *widget.ProgressBar {
	loading := widget.NewProgressBar()
	loading.Min = 0
	loading.Max = 100
	go loadAndShow(loading, compatible, navigationButtons)
	return loading
}

func loadAndShow(loading *widget.ProgressBar, compatible bool, navigationButtons *fyne.Container) {
	fyne.Do(func() {
		time.Sleep(10 * time.Millisecond)
		for i := 1.0; i <= 100; i += 1.0 {
			loading.SetValue(i)
			time.Sleep(loadingSpeed)
		}

		descL := widget.NewLabel("Compatibility results: ")
		label := getResultLabel(compatible)

		cfg.Window.SetContent(container.NewVBox(
			descL,
			label,
			layout.NewSpacer(),
			widget.NewSeparator(),
			widget.NewButton("See interesting information", rickrollOrRestore()),
			navigationButtons,
		))
	})

}

func getResultLabel(compatible bool) *widget.Label {
	var includeExtra bool
	label := widget.NewLabel("")
	nilPlayer := Player{}

	if player != nilPlayer {
		includeExtra = true
	}

	if compatible {
		label.SetText("You two are compatible, at least mathematically...")
		if includeExtra {
			if player.inRelationship || player.isMarried {
				label.SetText(label.Text + "\nAre you happy?\nIf yes, wait till you see drama.\nAbsolute cinema.")
			}
			if player.isSingle {
				label.SetText(label.Text + "\nStill pretty sure you'll get dumped.\nAbsolute cinema.")
			}
		}
	} else {
		label.SetText("Your choices are same as your mind\nPathetic.")
		if includeExtra {
			if player.inRelationship || player.isMarried {
				label.SetText(label.Text + "\nStarting to question your life choices?\nHilarious \nIt's just the start, not the end...")
			}
			if player.isSingle {
				label.SetText(label.Text + "\nDon't worry you'll die as a virgin too")
			}
		}
	}
	label.Alignment = fyne.TextAlignCenter
	return label
}
