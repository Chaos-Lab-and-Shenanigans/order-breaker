package tappedfunctions

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

var i = 0

func TapFuckJai(w *widget.Label) func() {
	return func() {
		i++
		text := fmt.Sprintf("Smashed jai %v times successfully", i)
		w.SetText(text)
	}
}
