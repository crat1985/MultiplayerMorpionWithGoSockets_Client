package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var a fyne.App
var mainWindow fyne.Window

func getButtons() []*fyne.Container {
	var buttons []*fyne.Container
	for x := 0; x < 9; x++ {
		btn1, color1, finalButton1 := getButtonWithColor(color.RGBA{R: 0, G: 0, B: 255})
		btn1.SetText("test")
		color1.FillColor = color.RGBA{R: 255}
		buttons = append(buttons, finalButton1)
	}
	return buttons
}

func undefined() {
	buttons := getButtons()
	morpioncontainer := container.NewGridWithRows(3, buttons[0], buttons[1], buttons[2], buttons[3], buttons[4], buttons[5], buttons[6], buttons[7], buttons[8])
	text := canvas.NewText("Morpion", color.RGBA{R: 255, B: 255})
	text.Alignment = fyne.TextAlignCenter
	maincontainer := container.NewBorder(text, nil, nil, nil, morpioncontainer)
	log.Print(maincontainer)
}

func login() {
	mainWindow.Hide()
	loginWin := a.NewWindow("Se connecter")
	pseudoEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	form := widget.NewForm(
		widget.NewFormItem("Pseudo : ", pseudoEntry),
		widget.NewFormItem("Mot de passe : ", passwordEntry),
	)
	form.OnSubmit = func() {
		log.Println("Submited !")
	}
	form.OnCancel = func() {
		log.Println("Canceled !")
	}
	loginWin.SetOnClosed(func() {
		mainWindow.Show()
	})
	loginWin.SetContent(form)
	loginWin.Resize(fyne.NewSize(350, loginWin.Canvas().Size().Height))
	loginWin.SetFixedSize(true)
	loginWin.CenterOnScreen()
	loginWin.Show()
}

func main() {
	a = app.NewWithID("Morpion Multijoueur en Go")
	mainWindow = a.NewWindow("Morpion multijoueur")
	button := widget.NewButton("Se connecter", login)
	mainWindow.SetContent(button)
	mainWindow.SetFixedSize(true)
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}

func getButtonWithColor(c color.Color) (*widget.Button, *canvas.Rectangle, *fyne.Container) {
	btn := widget.NewButton("", nil)
	btn_color := canvas.NewRectangle(c)
	btn_container := container.NewMax(btn_color, btn)
	return btn, btn_color, btn_container
}
