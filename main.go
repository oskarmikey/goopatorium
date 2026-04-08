package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow(" Button testi")

	fmt.Println("this shit works!")

	input := widget.NewEntry()
	input.SetPlaceHolder("yolo")

	input_saveBtnSize := fyne.NewSize(200, 50)
	saveBtn := widget.NewButton("save", func() {
		log.Println("saved info is: ", input.Text)
	})
	sizedSaveBtn := container.NewGridWrap(input_saveBtnSize, saveBtn)

	boopBtn := widget.NewButton("boops here", func() {
		log.Println("Boop!")
	})

	check := widget.NewCheck("Optional ?", func(value bool) {
		log.Println("chek isorwas set to", value)
	})

	radio := widget.NewRadioGroup([]string{"option1", "option2"}, func(value string) {
		log.Println("radio set to", value)
	})

	combo := widget.NewSelect([]string{"option3", "option4"}, func(value string) {
		log.Println("It was set to", value)
	})

	mainLayout := container.NewVBox(
		input,
		sizedSaveBtn,
		boopBtn,
		check,
		radio,
		combo,
	)

	myWindow.CenterOnScreen()

	myWindow.Canvas().Size()

	myWindow.SetContent(mainLayout)
	myWindow.ShowAndRun()

}
