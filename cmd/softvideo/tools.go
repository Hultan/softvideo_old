package main

import (
	"io/ioutil"
	"log"
	"path"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func builderGetObject(builder *gtk.Builder, name string) glib.IObject {
	obj, err := builder.GetObject(name)
	if err != nil {
		panic(err)
	}

	return obj
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

//
// func getExecutablePath() string {
//	fileName, _ := os.Executable()
//	return path.Dir(fileName)
// }

func filePathWalkDir(root string) ([]string, error) {
	var f []string
	// err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	//	if !info.IsDir() {
	//		f = append(f, path)
	//	}
	//	return nil
	// })
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		f = append(f, path.Join(root, file.Name()))
		// fmt.Println(file.Name(), file.IsDir())
	}
	return f, err
}
