package handleastrology

import "fyne.io/fyne/v2"

func StartAstro(w fyne.Window, logsCh chan string) func() {
	return func() {
		firstPage(w)
		logsCh <- "Not available yet"
	}
}
