package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func llmChat(a fyne.App) {
	w := a.NewWindow("GoopChat")

	// --- 1. THE SIDEBAR (LEFT) ---
	sideButtons := container.NewVBox(
		widget.NewButton("New Chat", func() { fmt.Println("New Chat logic") }),
		widget.NewSeparator(),
		widget.NewLabel("GPTs"),
		widget.NewButton("Source Engine Helper", func() {}),
		widget.NewButton("Fabric Modding Helper", func() {}),
	)
	sidebar := container.NewVScroll(sideButtons)

	// We wrap the sidebar in a container with a background or specific width
	// For now, let's keep it simple.
	sidebarArea := container.NewPadded(sidebar)

	// --- 2. THE CHAT DISPLAY (CENTER) ---
	chatDisplay := widget.NewRichTextFromMarkdown("# Where should we begin?")
	scrollContainer := container.NewVScroll(chatDisplay)

	// --- 3. THE INPUT AREA (BOTTOM) ---
	input := widget.NewEntry()
	input.SetPlaceHolder("Ask anything...")

	var fullHistory []Message

	// Defined FIRST so the button can see it
	sendFunc := func() {
		userText := input.Text
		if userText == "" {
			return
		}

		input.SetText("")
		fullHistory = append(fullHistory, Message{Role: "user", Content: userText})

		chatDisplay.ParseMarkdown("Thinking...")

		go func() {
			reply, err := getAIResponse(fullHistory)

			fyne.Do(func() {
				if err != nil {
					chatDisplay.ParseMarkdown("**Error:** " + err.Error())
					return
				}

				fullHistory = append(fullHistory, Message{Role: "assistant", Content: reply})

				// Show full history so it doesn't disappear!
				var historyLog string
				for _, m := range fullHistory {
					historyLog += fmt.Sprintf("**%s**: %s\n\n", m.Role, m.Content)
				}
				chatDisplay.ParseMarkdown(historyLog)
				scrollContainer.ScrollToBottom()
			})
		}()
	}

	sendBtn := widget.NewButton("Send", sendFunc)
	bottomRow := container.NewBorder(nil, nil, nil, sendBtn, input)

	// --- 4. THE LAYOUT LOGIC (NESTING) ---

	// First, put the scroll area and the input together for the "Right Side"
	chatArea := container.NewBorder(nil, bottomRow, nil, nil, scrollContainer)

	// Finally, put the sidebar on the left and the chatArea in the center
	// Center takes up all remaining space!
	content := container.NewBorder(nil, nil, sidebarArea, nil, chatArea)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600)) // Wider like a real chat app
	w.CenterOnScreen()
	w.Show()
}
