package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Morpion multijoueur")
	w.Resize(fyne.NewSize(720, 480))
	btn1, color1, finalButton1 := getButtonWithColor(color.RGBA{R: 0, G: 0, B: 255})
	btn1.SetText("test")
	color1.FillColor = color.RGBA{R: 255}
	// btn.ExtendBaseWidget(btn)
	button2 := widget.NewButton("", nil)
	button3 := widget.NewButton("", nil)
	button4 := widget.NewButton("", nil)
	button5 := widget.NewButton("", nil)
	button6 := widget.NewButton("", nil)
	button7 := widget.NewButton("", nil)
	button8 := widget.NewButton("", nil)
	button9 := widget.NewButton("", nil)
	morpioncontainer := container.NewGridWithRows(3, finalButton1, button2, button3, button4, button5, button6, button7, button8, button9)
	text := canvas.NewText("Morpion", color.White)
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
