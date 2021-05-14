package main

import (
	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"path"
	"path/filepath"
)

func FilePathWalkDir(root string) ([]string, error) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func builderGetObject(builder *gtk.Builder, name string) glib.IObject {
	obj, err := builder.GetObject(name)
	assertErr(err)

	return obj
}

func assertErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func assertConv(ok bool) {
	if !ok {
		log.Panic("invalid widget conversion")
	}
}

func playerReleaseMedia(player *vlc.Player) {
	player.Stop()
	if media, _ := player.Media(); media != nil {
		media.Release()
	}
}

func getExecutablePath() string {
	fileName, _ := os.Executable()
	return path.Dir(fileName)
}