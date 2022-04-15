package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "se.softteam.softvideo"

var player *vlc.Player
var playButton *gtk.ToolButton
var builder *gtk.Builder
var playerWindow *gdk.Window

// var nextButton *gtk.Button
var appWin *gtk.ApplicationWindow
var ok bool
var files []string
var paths *pathList
var recentMenu *gtk.MenuItem
var lastPath string
var slider *gtk.Scale
var ticker *time.Ticker
var ignoreTick bool

func main() {
	gtk.Init(nil)

	// Initialize libVLC module.
	err := vlc.Init("--quiet", "--no-xlib")
	if err != nil {
		panic(err)
	}

	// Set RNG seed
	rand.Seed(time.Now().UnixNano())

	// Create a new player.
	player, err = vlc.NewPlayer()
	if err != nil {
		panic(err)
	}

	// Create new GTK application.
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		panic(err)
	}

	app.Connect("activate", func() {
		// Load application layout.
		builder, err = gtk.BuilderNewFromFile("../assets/layout.glade")
		if err != nil {
			panic(err)
		}

		// Get application window.
		appWin, ok = builderGetObject(builder, "appWindow").(*gtk.ApplicationWindow)
		assertConv(ok)
		appWin.SetTitle(fmt.Sprintf("%s - %s", applicationName, applicationVersion))

		slider = builderGetObject(builder, "slider").(*gtk.Scale)

		// Get play button.
		playButton, ok = builderGetObject(builder, "toolbarPlayButton").(*gtk.ToolButton)
		assertConv(ok)

		// // Get next button.
		// nextButton, ok = builderGetObject(builder, "nextButton").(*gtk.Button)
		// assertConv(ok)

		// Get recent menu item.
		recentMenu, ok = builderGetObject(builder, "recentMenuItem").(*gtk.MenuItem)
		assertConv(ok)

		// Path list
		paths = newPathList()
		paths.load()
		populateRecentMenu()

		// Add builder signal handlers.
		signals := map[string]interface{}{
			"onRealizePlayerArea": func(playerArea *gtk.DrawingArea) {
				// Set window for the player.
				playerWindow, err = playerArea.GetWindow()
				if err != nil {
					panic(err)
				}

				err = setPlayerWindow(player, playerWindow)
				if err != nil {
					panic(err)
				}
			},
			"onDrawPlayerArea": func(playerArea *gtk.DrawingArea, cr *cairo.Context) {
				cr.SetSourceRGB(0, 0, 0)
				cr.Paint()
			},
			"onActivateOpenFile": onActivateOpenFile,
			"onActivateQuit": func() {
				app.Quit()
			},
			"onClickPlayButton":         onClickPlayButton,
			"onClickStopButton":         onClickStopButton,
			"onClickJumpButton":         onClickJumpButton,
			"onClickPreviousButton":     onClickPreviousButton,
			"onClickRewindButton":       onClickRewindButton,
			"onClickForwardButton":      onClickForwardButton,
			"onClickNextButton":         onClickNextButton,
			"onClickFullscreenButton":   onClickFullscreenButton,
			"onClickUnfullscreenButton": onClickUnfullscreenButton,
			"onSliderValueChanged":      onSliderValueChanged,
		}
		builder.ConnectSignals(signals)

		appWin.ShowAll()
		app.AddWindow(appWin)
	})

	// Cleanup on exit.
	app.Connect("shutdown", func() {
		paths.save()
		playerReleaseMedia(player)
		player.Release()
		vlc.Release()
	})

	// Launch the application.
	os.Exit(app.Run(os.Args))
}

func onSliderValueChanged(slider *gtk.Scale) {
	if ignoreTick {
		ignoreTick = false
		return
	}
	value := slider.GetValue()
	player.SetMediaPosition(float32(value))
}
