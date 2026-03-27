package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hellowwwwyy World")
	screenSize := fyne.CurrentApp().Driver().AllWindows()[0].Canvas().Size()

	screenWidth := screenSize.Width
	screenHeight := screenSize.Height

	targetWidth := screenWidth * 0.2
	targetHeight := screenHeight * 0.4

	if targetWidth < 300 {
		targetWidth = 300
	}
	if targetWidth > 800 {
		targetWidth = 800
	}

	myWindow.Resize(fyne.NewSize(targetWidth, targetHeight))
	myWindow.CenterOnScreen()

	myWindow.SetContent(widget.NewLabel("Hello World!"))
	myWindow.ShowAndRun()
	tidyup()
}

func tidyup() {
	fmt.Println("Testing123zzz")
}
