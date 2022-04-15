package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

func onActivateOpenFile() {
	fileDialog, err := gtk.FileChooserDialogNewWith2Buttons(
		"Choose file...",
		appWin,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		"Cancel", gtk.RESPONSE_DELETE_EVENT,
		"Open", gtk.RESPONSE_ACCEPT)
	if err != nil {
		panic(err)
	}
	defer fileDialog.Destroy()

	fileFilter, err := gtk.FileFilterNew()
	if err != nil {
		panic(err)
	}

	fileFilter.SetName("Media files")
	fileFilter.AddPattern("*.avi")
	fileFilter.AddPattern("*.flv")
	fileFilter.AddPattern("*.mkv")
	fileFilter.AddPattern("*.mpg")
	fileFilter.AddPattern("*.mp4")
	fileFilter.AddPattern("*.wmv")
	fileDialog.AddFilter(fileFilter)
	if lastPath != "" {
		fileDialog.SetCurrentFolder(lastPath)
	}

	if result := fileDialog.Run(); result == gtk.RESPONSE_ACCEPT {
		// Get selected filename.
		dirname, _ := fileDialog.GetCurrentFolder()

		changePath(dirname)
	}
}

func onClickPlayButton(playButton *gtk.ToolButton) {
	if media, _ := player.Media(); media == nil {
		return
	}

	if player.IsPlaying() {
		player.SetPause(true)
		playButton.SetLabel("gtk-media-play")
	} else {
		player.Play()
		playButton.SetLabel("gtk-media-pause")
	}
}

func onClickJumpButton(nextButton *gtk.ToolButton) {
	// Stop playing the current video
	onClickStopButton(nil)

	if files == nil {
		return
	}
	v := rand.Intn(len(files)) // range is min to max
	filename := files[v]

	// Load media and start playback.
	if _, err := player.LoadMediaFromPath(filename); err != nil {
		log.Printf("Cannot load selected media: %s\n", err)
		return
	}

	player.Play()
	// player.SetMediaPosition(0.2)
	playButton.SetLabel("gtk-media-pause")
}

func onClickStopButton(_ *gtk.ToolButton) {
	ticker.Stop()
	player.Stop()
	playButton.SetLabel("gtk-media-play")
}

func onClickPreviousButton(_ *gtk.ToolButton) {
	pos, _ := player.MediaPosition()
	newPos := pos - 0.05
	if newPos <= 0 {
		newPos = 0.01
	}
	player.SetMediaPosition(newPos)
}

func onClickRewindButton(_ *gtk.ToolButton) {
	pos, _ := player.MediaPosition()
	newPos := pos - 0.005
	if newPos <= 0 {
		newPos = 0.01
	}
	player.SetMediaPosition(newPos)
}

func onClickForwardButton(_ *gtk.ToolButton) {
	pos, _ := player.MediaPosition()
	newPos := pos + 0.005
	if newPos <= 0 {
		newPos = 0.01
	}
	player.SetMediaPosition(newPos)
}

func onClickNextButton(_ *gtk.ToolButton) {
	pos, _ := player.MediaPosition()
	newPos := pos + 0.05
	if newPos > 1 {
		newPos = 0.99
	}
	player.SetMediaPosition(newPos)
}

func changePath(path string) {
	// Release current media, if any.
	playerReleaseMedia(player)

	// filepath.Walk
	f, err := filePathWalkDir(path)
	if err != nil {
		panic(err)
	}

	files = f
	lastPath = path
	paths.addPath(path)
	populateRecentMenu()

	v := rand.Intn(len(files)) // range is min to max
	filename := files[v]

	// Load media and start playback.
	if _, err := player.LoadMediaFromPath(filename); err != nil {
		log.Printf("Cannot load selected media: %s\n", err)
		return
	}

	ticker = time.NewTicker(2 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				value, err := player.MediaPosition()
				if err != nil {
					return
				}
				ignoreTick = true
				slider.SetValue(float64(value))
			}
		}
	}()

	player.Play()
	// player.SetMediaPosition(0.2)
	playButton.SetLabel("gtk-media-pause")
}

func onClickFullscreenButton() {
	appWin.Fullscreen()
}

func onClickUnfullscreenButton() {
	appWin.Unfullscreen()
}
