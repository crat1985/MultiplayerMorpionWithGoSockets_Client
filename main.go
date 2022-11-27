package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var a fyne.App
var mainWindow fyne.Window
var conn net.Conn
var pseudoEntry *widget.Entry
var addressOfServer *widget.Entry
var portOfServer *widget.Entry
var isCurrentlyLogin bool
var joinPartyEntry *widget.Entry

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

func PartyCreated() {
	dialog.NewInformation("Partie créée", "Partie créée avec succès !", mainWindow).Show()
}

func CreateParty() {
	_, err := conn.Write([]byte("createparty"))
	if err != nil {
		dialog.NewError(err, mainWindow).Show()
		return
	}
	slice := make([]byte, 1024)
	n, err := conn.Read(slice)
	if err != nil {
		dialog.NewError(err, mainWindow).Show()
		return
	}
	response := string(slice[:n])
	if response == "partycreated" {
		PartyCreated()
		return
	}
}

func JoinParty() {
	if joinPartyEntry.Text == "" {
		dialog.NewError(errors.New("ID de la partie nécessaire"), mainWindow).Show()
		return
	}
	log.Println("En développement...")
}

func LoginSuccessfully() {
	createPartyButton := widget.NewButton("Créer une partie", CreateParty)
	orLabel := widget.NewLabel("OU")
	orLabel.Alignment = fyne.TextAlignCenter
	joinPartyEntry = widget.NewEntry()
	joinPartyButton := widget.NewButton("Rejoindre une partie", JoinParty)
	joinPartyForm := widget.NewForm(
		widget.NewFormItem("ID de la partie : ", joinPartyEntry),
	)
	mainContainer := container.NewVBox(createPartyButton, orLabel, joinPartyForm, joinPartyButton)
	mainWindow.SetContent(mainContainer)
}

func SendPseudo() error {
	conn.Write([]byte(pseudoEntry.Text))
	isPseudoOk := make([]byte, 1024)
	n, err := conn.Read(isPseudoOk)
	if err != nil {
		dialog.NewError(err, mainWindow).Show()
		conn.Close()
		return errors.New("erreur lors de la lecture de la réponde du serveur")
	}
	isPseudoOkString := string(isPseudoOk[:n])
	if isPseudoOkString != "pseudook" {
		conn.Close()
		return errors.New(isPseudoOkString)
	}
	fmt.Println("Connecté !")
	return nil
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
		dialog.NewError(err, mainWindow).Show()
		return
	}
	conn, err = net.Dial("tcp", address+":"+port)
	if err != nil {
		dialog.NewError(err, mainWindow).Show()
		return
	}
	err = SendPseudo()
	if err != nil {
		log.Print(err)
		return
	}
	LoginSuccessfully()
}

func main() {
	a = app.NewWithID("Morpion Multijoueur en Go")
	mainWindow = a.NewWindow("Morpion multijoueur")
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
	mainWindow.SetCloseIntercept(
		dialog.NewConfirm("Quitter ?", "Voulez-vous vraiment quitter le jeu ?", func(b bool) {
			if b {
				mainWindow.Close()
			}
		}, mainWindow).Show,
	)
	mainWindow.SetContent(form)
	mainWindow.Resize(fyne.NewSize(400, mainWindow.Canvas().Size().Height))
	mainWindow.SetFixedSize(true)
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}
