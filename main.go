
package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Login")
	myWindow.Resize(fyne.NewSize(400, 300))

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	statusLabel := widget.NewLabel("")
	statusLabel.Alignment = fyne.TextAlignCenter

	loginButton := widget.NewButton("Login", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text

		if username == "admin" && password == "admin" {
			statusLabel.SetText("Login Successful!")
			statusLabel.Refresh()
			log.Println("Login Successful!")
		} else {
			statusLabel.SetText("Invalid Credentials")
			statusLabel.Refresh()
			log.Println("Invalid Credentials")
		}
	})

	cancelButton := widget.NewButton("Cancel", func() {
		usernameEntry.SetText("")
		passwordEntry.SetText("")
		statusLabel.SetText("Login Cancelled")
		statusLabel.Refresh()
		log.Println("Login Cancelled")
	})

	// Create a background rectangle for visual appeal
	background := canvas.NewRectangle(color.RGBA{R: 240, G: 240, B: 240, A: 255})
	background.SetMinSize(fyne.NewSize(400, 300))

	// Layout the widgets
	content := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Username:"),
		usernameEntry,
		widget.NewLabel("Password:"),
		passwordEntry,
		container.New(layout.NewHBoxLayout(),
			layout.NewSpacer(),
			loginButton,
			cancelButton,
			layout.NewSpacer(),
		),
		statusLabel,
	)

	// Center the content in the window
	centeredContent := container.New(layout.NewCenterLayout(), content)

	// Overlay the content on the background
	myWindow.SetContent(container.New(layout.NewMaxLayout(), background, centeredContent))

	myWindow.ShowAndRun()
}
