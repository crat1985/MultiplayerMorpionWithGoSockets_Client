package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var a fyne.App
var loginWin fyne.Window
var conn net.Conn
var pseudoEntry *widget.Entry
var addressOfServer *widget.Entry
var portOfServer *widget.Entry
var isCurrentlyLogin bool

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

func getInfos() (address, port, pseudo string, err error) {
	if pseudoEntry.Text == "" {
		return "", "", "", errors.New("veuillez entrer un pseudo")
	}
	if addressOfServer.Text == "" {
		address = addressOfServer.PlaceHolder
	}
	if portOfServer.Text == "" {
		port = portOfServer.PlaceHolder
	}
	return address, port, pseudo, nil
}

func SendPseudo() {
	conn.Write([]byte(pseudoEntry.Text))
	isPseudoOk := make([]byte, 1024)
	n, err := conn.Read(isPseudoOk)
	if err != nil {
		dialog.NewError(err, loginWin).Show()
		return
	}
	isPseudoOkString := string(isPseudoOk[:n])
	if isPseudoOkString != "pseudook" {
		log.Println(isPseudoOkString)
		conn.Close()
		return
	}
	fmt.Println("Connecté !")
}

func login() {
	if isCurrentlyLogin {
		return
	}
	defer func() {
		isCurrentlyLogin = false
	}()
	var err error
	address, port, _, err := getInfos()
	if err != nil {
		dialog.NewError(err, loginWin).Show()
		return
	}
	conn, err = net.Dial("tcp", address+":"+port)
	if err != nil {
		dialog.NewError(err, loginWin).Show()
		return
	}
	SendPseudo()
}

func main() {
	a = app.NewWithID("Morpion Multijoueur en Go")
	loginWin = a.NewWindow("Morpion multijoueur")
	pseudoEntry = widget.NewEntry()
	pseudoEntry.SetPlaceHolder("Votre pseudo ici")
	addressOfServer = widget.NewEntry()
	addressOfServer.SetPlaceHolder("localhost")
	portOfServer = widget.NewEntry()
	portOfServer.SetPlaceHolder("8888")
	form := widget.NewForm(
		widget.NewFormItem("Pseudo : ", pseudoEntry),
		widget.NewFormItem("Adresse du serveur : ", addressOfServer),
		widget.NewFormItem("Port du serveur : ", portOfServer),
	)
	form.OnSubmit = login
	form.OnCancel = a.Quit
	loginWin.SetCloseIntercept(
		dialog.NewConfirm("Quitter ?", "Voulez-vous vraiment quitter le jeu ?", func(b bool) {
			if b {
				loginWin.Close()
			}
		}, loginWin).Show,
	)
	loginWin.SetContent(form)
	loginWin.Resize(fyne.NewSize(400, loginWin.Canvas().Size().Height))
	loginWin.SetFixedSize(true)
	loginWin.CenterOnScreen()
	loginWin.ShowAndRun()
}

// Inutile pour l'instant
func getButtonWithColor(c color.Color) (*widget.Button, *canvas.Rectangle, *fyne.Container) {
	btn := widget.NewButton("", nil)
	btn_color := canvas.NewRectangle(c)
	btn_container := container.NewMax(btn_color, btn)
	return btn, btn_color, btn_container
}
