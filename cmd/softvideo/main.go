package main

import (
	"math/rand"
	"os"
	"path"
	"time"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "com.github.libvlc-go.gtk3-media-player-example"

var player *vlc.Player
var playButton *gtk.Button
var nextButton *gtk.Button
var appWin *gtk.ApplicationWindow
var ok bool
var files []string
var pathList *PathList
var recentMenu *gtk.MenuItem
var lastPath string

func main() {
	// Initialize libVLC module.
	err := vlc.Init("--quiet", "--no-xlib")
	assertErr(err)

	// Set RNG seed
	rand.Seed(time.Now().UnixNano())

	// Path list
	pathList = NewPathList()

	// Create a new player.
	player, err = vlc.NewPlayer()
	assertErr(err)

	// Create new GTK application.
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	assertErr(err)

	app.Connect("activate", func() {
		// Load application layout.
		builder, err := gtk.BuilderNewFromFile(path.Join(getExecutablePath(),"layout.glade"))
		assertErr(err)

		// Get application window.
		appWin, ok = builderGetObject(builder, "appWindow").(*gtk.ApplicationWindow)
		assertConv(ok)

		// Get play button.
		playButton, ok = builderGetObject(builder, "playButton").(*gtk.Button)
		assertConv(ok)

		// Get next button.
		nextButton, ok = builderGetObject(builder, "nextButton").(*gtk.Button)
		assertConv(ok)

		// Get recent menu item.
		recentMenu, ok = builderGetObject(builder, "recentMenuItem").(*gtk.MenuItem)
		assertConv(ok)

		// Add builder signal handlers.
		signals := map[string]interface{}{
			"onRealizePlayerArea": func(playerArea *gtk.DrawingArea) {
				// Set window for the player.
				playerWindow, err := playerArea.GetWindow()
				assertErr(err)
				err = setPlayerWindow(player, playerWindow)
				assertErr(err)
			},
			"onDrawPlayerArea": func(playerArea *gtk.DrawingArea, cr *cairo.Context) {
				cr.SetSourceRGB(0, 0, 0)
				cr.Paint()
			},
			"onActivateOpenFile": onActivateOpenFile,
			"onActivateQuit": func() {
				app.Quit()
			},
			"onClickPlayButton": onClickPlayButton,
			"onClickStopButton": onClickStopButton,
			"onClickJumpButton": onClickJumpButton,
			"onClickPreviousButton": onClickPreviousButton,
			"onClickRewindButton": onClickRewindButton,
			"onClickForwardButton": onClickForwardButton,
			"onClickNextButton": onClickNextButton,
		}
		builder.ConnectSignals(signals)

		appWin.ShowAll()
		app.AddWindow(appWin)
	})

	// Cleanup on exit.
	app.Connect("shutdown", func() {
		playerReleaseMedia(player)
		player.Release()
		vlc.Release()
	})

	// Launch the application.
	os.Exit(app.Run(os.Args))
}
