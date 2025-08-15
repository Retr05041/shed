package main

import (
	"image/color"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	MAX_WIDTH  = 1200 // 600
	MAX_HEIGHT = 800  //400

	BLACK      = color.NRGBA{86, 86, 86, 255}
	GREEN      = color.NRGBA{0, 153, 51, 255}
	BACKGROUND = color.NRGBA{255, 253, 208, 255}
)

// state provided by server
type gameState struct {
	serverConn net.Conn

	tableCode    int32
	turnOrder    []int32          // Turn order is the order of GUID's in this slice
	otherPlayers map[int32]string // otherPlayers will be a map of GUID to nickname
}

func main() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("Shed")

	// Init all scene elements
	initElements()

	// Run the first scene
	showMenuUI(mainWindow)

	// Run GUI
	mainWindow.Resize(fyne.NewSize(float32(MAX_WIDTH), float32(MAX_HEIGHT)))
	mainWindow.ShowAndRun()
}
