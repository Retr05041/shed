package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func setWindowContent(window fyne.Window, content fyne.CanvasObject) {
	if window.Content() != content {
		window.SetContent(content)
		window.Resize(fyne.NewSize(float32(MAX_WIDTH), float32(MAX_HEIGHT))) // Ensure size consistency
	}
}

// Main Menu
func showMenuUI(givenWindow fyne.Window) {
	banner := canvas.NewText("SHED", GREEN)
	banner.TextSize = 64
	banner.TextStyle = fyne.TextStyle{Bold: true, Italic: false}
	banner.Alignment = fyne.TextAlignCenter

	or := canvas.NewText("Or", BLACK)
	or.TextSize = 18
	or.TextStyle = fyne.TextStyle{Bold: true, Italic: false}
	or.Alignment = fyne.TextAlignCenter

	nickname := widget.NewEntry()
	nickname.SetPlaceHolder("Nickname...")

	host := widget.NewButton("Host a table", func() {
	})

	connect := widget.NewButton("Join a table", func() {
	})

	background := canvas.NewRectangle(BACKGROUND)

	setWindowContent(givenWindow,
		container.NewStack(background, container.NewCenter(
			container.NewGridWrap(
				fyne.NewSize(float32(MAX_WIDTH)/3, float32(MAX_HEIGHT)/3),
				container.NewCenter(banner),
				container.NewVBox(
					nickname,
					container.NewGridWithColumns(3, host, or, connect))))))
}

// Host UI is the same as connectedUI but without settings
// func showHostUI(givenWindow fyne.Window) {
// 	playButton := widget.NewButton("Play", func() {
// 		// channelmanager.FGUI_ActionChan <- channelmanager.ActionType{Action: "startRound"}
// 	})
// 	copyTableCode := widget.NewButton("Copy table code", func() {
// 		givenWindow.Clipboard().SetContent(string(state.tableCode))
// 	})

// 	setWindowContent(givenWindow,
// 		container.NewCenter(
// 			container.NewVBox(
// 				numOfPlayers,
// 				container.NewHBox(copyTableCode),
// 				playButton)))
// }

// showConnectedUI shows the lobby while the host waits for everyone to join - might want to remove and just wait for the host to click start in the game screen?
// func showConnectedUI(givenWindow fyne.Window) {
// 	waiting := widget.NewLabel("Waiting for host to begin game!")
// 	setWindowContent(givenWindow,
// 		container.NewCenter(
// 			container.NewVBox(numOfPlayers, waiting)))
// }

// Main game screen
// func showGameScreen(givenWindow fyne.Window) {
// setWindowContent(givenWindow,
// container.NewBorder(
// nil,
// nil,
// playerCards,
// container.NewCenter(
// container.NewVBox(
// container.NewCenter(potLabel),
// boardGrid,
// container.NewCenter(
// container.NewHBox(
// handGrid,
// container.NewVBox(
// foldButton,
// callButton,
// container.NewHBox(raiseButton, valueLabel),
// betSlider),
// checkButton))))))
// }

// Helper loading screen while server responds
func showLoadingScreen(givenWindow fyne.Window) {
	loadingText := canvas.NewText("Loading...", GREEN)
	loadingText.TextSize = 24
	loadingText.Alignment = fyne.TextAlignCenter

	progressBar := widget.NewProgressBarInfinite()
	progressBar.Start()

	loadingContainer := container.NewCenter(
		container.NewVBox(
			container.NewCenter(loadingText),
			container.NewCenter(progressBar),
		),
	)

	setWindowContent(givenWindow, loadingContainer)
}
