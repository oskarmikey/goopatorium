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
	myWindow := myApp.NewWindow(" Buggati power engine")

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

	boopBtnSpawnPopup := widget.NewButton("boops here for popup", func() {
		spawnPopup(myApp, "Suprise thy selves", "I am thus descented from thee thus i serve you", true)
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

	openCat := widget.NewButton("Open seasamy", func() {
		log.Println("Boop!")
		tgaillmchat(myApp)
	})

	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Ouchie", Widget: entry}},
		OnSubmit: func() {
			log.Println("Form submite", entry.Text)
			log.Println("multline", textArea.Text)
		},
	}

	form.Append("text", textArea)

	mainLayout := container.NewVBox(
		input,
		sizedSaveBtn,
		boopBtn,
		boopBtnSpawnPopup,
		openCat,
		check,
		radio,
		combo,
		form,
	)

	myWindow.CenterOnScreen()

	myWindow.Canvas().Size()

	myWindow.SetContent(mainLayout)
	myWindow.ShowAndRun()

}

func summonTheePopup(a fyne.App, title string, message string) {
	win := a.NewWindow(title)

	label := widget.NewLabel(message)
	closeBtn := widget.NewButton("Go away", func() {
		win.Close()
	})

	closeBtn2 := widget.NewButton("Okay", func() {
		win.Close()
	})

	win.SetContent(container.NewVBox(
		label,
		closeBtn,
		closeBtn2,
	))
	win.CenterOnScreen()
	win.Resize(fyne.NewSize(200, 100))
	win.Show()
}

func spawnPopup(a fyne.App, title string, message string, isthereachoice bool) {
	win := a.NewWindow(title)

	label := widget.NewLabel(message)
	closeBtn := widget.NewButton("Go away", func() {
		win.Close()
	})

	var blueorredpill fyne.CanvasObject
	if isthereachoice {

		blueorredpill = container.NewVBox(
			widget.NewButton("of course", func() {
				log.Println("The sire has concluded an affermative choice to:", message)
				win.Close()
			}),
			widget.NewButton("i demand you be taken to the prision!", func() {
				log.Printf("The sire said no to: %s Were doomed now!", message)
				win.Close()
			}),
		)
	} else {
		blueorredpill = widget.NewButton("Yes", func() {
			win.Close()
		})
	}

	win.SetContent(container.NewVBox(
		label,
		blueorredpill,
		closeBtn,
	))
	win.Resize(fyne.NewSize(300, 150))
	win.CenterOnScreen()
	win.Show()
}

func tgaillmchat(a fyne.App) {
	cWin := a.NewWindow("Chat")

	label := widget.NewLabel("test")

	textArea := widget.
		cWin.SetContent(container.NewVBox(
		label,
	))

	cWin.CenterOnScreen()
	cWin.Resize(fyne.NewSize(200, 100))
	cWin.Show()
}
