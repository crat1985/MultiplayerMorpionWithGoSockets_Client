package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getButtons() []*fyne.Container {
	x := 0
	var buttons []*fyne.Container
	for x < 9 {
		btn1, color1, finalButton1 := getButtonWithColor(color.RGBA{R: 0, G: 0, B: 255})
		btn1.SetText("test")
		color1.FillColor = color.RGBA{R: 255}
		buttons = append(buttons, finalButton1)
		x++
	}
	return buttons
}

func main() {
	a := app.New()
	w := a.NewWindow("Morpion multijoueur")
	w.Resize(fyne.NewSize(720, 480))
	buttons := getButtons()
	morpioncontainer := container.NewGridWithRows(3, buttons[0], buttons[1], buttons[2], buttons[3], buttons[4], buttons[5], buttons[6], buttons[7], buttons[8])
	text := canvas.NewText("Morpion", color.RGBA{R: 255, B: 255})
	text.Alignment = fyne.TextAlignCenter
	maincontainer := container.NewBorder(text, nil, nil, nil, morpioncontainer)
	w.SetContent(maincontainer)
	w.ShowAndRun()
}

func getButtonWithColor(c color.Color) (*widget.Button, *canvas.Rectangle, *fyne.Container) {
	btn := widget.NewButton("", nil)
	btn_color := canvas.NewRectangle(c)
	btn_container := container.NewMax(btn_color, btn)
	return btn, btn_color, btn_container
}
