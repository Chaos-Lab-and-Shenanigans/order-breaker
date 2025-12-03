package astrology

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Chaos-Lab-and-Shenanigans/order-breaker/internal/config"
)

// This is a mess, needs refactoring
var (
	compatibilityLoadingSpeed = 20 * time.Millisecond
	compatible                bool
)

func StartCompatibilityChecker() func() {
	return func() {
		initNav()
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

	config.Cfg.Window.SetContent(mainContainer)
}

func calcStatus(p1Dob *widget.DateEntry, p2Dob *widget.DateEntry) func() {
	return func() {
		t1 := p1Dob.Date
		t2 := p2Dob.Date

		if t1 == nil || t2 == nil {
			return
		}

		config.User.Dob = p1Dob.Text
		y1, m1, d1 := t1.Date()
		y2, m2, d2 := t2.Date()

		//random shit to determine compatibility, even I didn't bother understanding it
		if (y1-3 < y2 && y2 < y1+3) && (m1-2 < m2 && m2 < m1+2) && (d1-5 < d2 && d2 < d1+5) {
			compatible = true
		}

		showCompatibility()
	}
}

func showCompatibility() {
	descL := widget.NewLabel("Checking compatibility...")
	loading := getLoading()

	config.Cfg.Window.SetContent(container.NewVBox(
		descL,
		layout.NewSpacer(),
		loading,
		widget.NewSeparator(),
		config.NavigationButtons,
	))
}

func getLoading() *widget.ProgressBar {
	loading := widget.NewProgressBar()
	go startLoading(loading)
	return loading
}

func startLoading(loading *widget.ProgressBar) {
	time.Sleep(10 * time.Millisecond)
	for i := 0.0; i <= 1.0; i += 0.1 {
		fyne.Do(func() { loading.SetValue(i) })
		time.Sleep(compatibilityLoadingSpeed)
	}

	fyne.Do(setThirdPageWindow)
}

func setThirdPageWindow() {
	descL := widget.NewLabel("Compatibility results: ")
	label := getCompatibilityLabel(compatible)

	if check() {
		config.Cfg.Window.SetContent(container.NewVBox(
			descL,
			label,
			layout.NewSpacer(),
			widget.NewSeparator(),
			widget.NewButton("See guessed personality", resultPage()),
			config.NavigationButtons,
		))
	} else {
		config.Cfg.Window.SetContent(container.NewVBox(
			descL,
			label,
			layout.NewSpacer(),
			widget.NewSeparator(),
			widget.NewButton("See interesting information", func() { rickrollOrRestore() }),
			config.NavigationButtons,
		))
	}
}

func getCompatibilityLabel(compatible bool) *widget.Label {
	var includeExtra bool
	label := widget.NewLabel("")

	if check() {
		includeExtra = true
	}

	if compatible {
		label.SetText("You two are compatible, at least mathematically...")
		if includeExtra {
			if config.User.Status == config.InRelationship || config.User.Status == config.IsMarried {
				label.SetText(label.Text + "\nYou finally got what everyone wants. \nDon’t worry, you’ll ruin it like everything else.")
			}
			if config.User.Status == config.IsSingle {
				label.SetText(label.Text + "\nEven when the universe sets you up to win, you still end up alone. \nThat’s not bad luck — that’s you.")
			}
		}
	} else {
		label.SetText("Your choices are same as your mind.\nPathetic.")
		if includeExtra {
			if config.User.Status == config.InRelationship || config.User.Status == config.IsMarried {
				label.SetText(label.Text + "\nThis relationship is two bad decisions dating each other. \nIt won’t last — but the regret will.")
			}
			if config.User.Status == config.IsSingle {
				label.SetText(label.Text + "\nNot only are you alone — you’re meant to be alone. \nDestiny itself said ‘nah.")
			}
		}
	}
	label.Alignment = fyne.TextAlignCenter
	return label
}
