package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Password Generator")

	w.SetContent(Container())

	w.Resize(fyne.NewSize(550, 100))
	w.SetFixedSize(true)
	w.SetPadded(true)

	w.ShowAndRun()
	myApp.Quit()
}

func handleQuit(a fyne.App) {
	a.Quit()
}
