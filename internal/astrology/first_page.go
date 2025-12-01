package astrology

import (
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	inputDate = widget.NewDateEntry()
)

func firstPage() {
	getAgeL := widget.NewLabel("Enter your date of birth:")
	inputDate.SetPlaceHolder("MM/DD/YYYY")

	page := container.NewVBox(
		getAgeL,
		inputDate,
		widget.NewSeparator(),
		layout.NewSpacer(),
		widget.NewButton("Next", showDetails()),
	)

	cfg.Window.SetContent(page)
}

func showDetails() func() {
	return func() {
		w := cfg.Window
		err := inputDate.Validate()
		if err != nil {
			inputDate.SetText("")
			inputDate.SetPlaceHolder("Invalid date(Valid format: MM/DD/YYYY)")
			return
		}

		player.dob = inputDate.Text
		descZodiacL := widget.NewLabel("Your zodiac sign:")

		zodiac := GetZodiacSign(*inputDate.Date)
		if player.zodiacSign == nil {
			player.zodiacSign = widget.NewLabel(zodiac)
		} else {
			player.zodiacSign.SetText(zodiac)
		}
		nextB := widget.NewButton("Continue", next())

		page := container.New(
			layout.NewGridLayout(2),
			descZodiacL,
			player.zodiacSign,
		)

		w.SetContent(container.NewVBox(
			page,
			layout.NewSpacer(),
			widget.NewSeparator(),
			nextB,
		))
	}
}

func next() func() {
	return func() {
		secondPage()
	}
}

func GetZodiacSign(date time.Time) string {
	month := date.Month()
	day := date.Day()

	switch {
	case (month == time.March && day >= 21) || (month == time.April && day <= 19):
		return "Aries"
	case (month == time.April && day >= 20) || (month == time.May && day <= 20):
		return "Taurus"
	case (month == time.May && day >= 21) || (month == time.June && day <= 20):
		return "Gemini"
	case (month == time.June && day >= 21) || (month == time.July && day <= 22):
		return "Cancer"
	case (month == time.July && day >= 23) || (month == time.August && day <= 22):
		return "Leo"
	case (month == time.August && day >= 23) || (month == time.September && day <= 22):
		return "Virgo"
	case (month == time.September && day >= 23) || (month == time.October && day <= 22):
		return "Libra"
	case (month == time.October && day >= 23) || (month == time.November && day <= 21):
		return "Scorpio"
	case (month == time.November && day >= 22) || (month == time.December && day <= 21):
		return "Sagittarius"
	case (month == time.December && day >= 22) || (month == time.January && day <= 19):
		return "Capricorn"
	case (month == time.January && day >= 20) || (month == time.February && day <= 18):
		return "Aquarius"
	case (month == time.February && day >= 19) || (month == time.March && day <= 20):
		return "Pisces"
	}

	return "Unknown"
}
