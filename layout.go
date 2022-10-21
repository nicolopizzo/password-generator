package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"passwordGenerator/algo"
	"strconv"
)

var charOptions = []string{
	"Cifre",
	"Maiuscola",
	"Caratteri speciali",
}

var lengthOptions = []string{"8", "16", "24"}

var chosenOptions []string
var L int

var translator = map[string]string{
	"Cifre":              "digits",
	"Maiuscola":          "maiusc",
	"Caratteri speciali": "special",
}

func Container() *fyne.Container {
	formChars := container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel("Scegli i caratteri inseriti"),
		widget.NewCheckGroup(charOptions, handleChangeChars),
	)

	radioGroup := widget.NewRadioGroup(lengthOptions, handleChangeLength)
	radioGroup.SetSelected(lengthOptions[1])
	radioGroup.Required = true

	label2 := widget.NewLabel("Scegli la lunghezza")
	choiceLength := container.New(layout.NewVBoxLayout(), label2, radioGroup)

	passwordEntry := widget.NewEntry()
	passwordEntry.Text = "Password"
	passwordEntry.Disable()

	passwordButton := widget.NewButton("Genera Password", func() { handleButtonClick(passwordEntry) })
	copyButton := widget.NewButton("Copia Password", func() { handleCopyPassword(passwordEntry.Text) })

	c := container.New(
		layout.NewVBoxLayout(),
		container.New(layout.NewGridLayout(2), formChars, choiceLength),
		container.New(layout.NewGridLayout(2), passwordEntry, copyButton),
		passwordButton,
	)

	return c
}


func handleCopyPassword(password string) {
	clipboard.WriteAll(password)
}

func handleButtonClick(w *widget.Entry) {
	parameters := make(map[string]bool)
	for _, k := range chosenOptions {
		parameters[k] = true
	}

	password := algo.NewPassword(parameters, L)
	w.Enable()
	w.Text = password
	w.Disable()
}

func translate(s string) string {
	return translator[s]
}

func handleChangeChars(choices []string) {
	chosenOptions = Map(translate, choices)
}

func handleChangeLength(s string) {
	L, _ = strconv.Atoi(s)
}
