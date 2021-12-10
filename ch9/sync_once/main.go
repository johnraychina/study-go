package sync_once

import (
	"image"
	"os"
	"sync"
)

func main() {

}

var loadIconOnce sync.Once
var icons map[string]image.Image

func Icon(name string) image.Image {
	loadIconOnce.Do(loadIcons)
	return icons[name]
}

func loadIcons() {
	icons = make(map[string]image.Image)
	icons["spades.png"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("hearts.png")
	icons["diamonds.png"] = loadIcon("diamonds.png")
	icons["clubs.png"] = loadIcon("clubs.png")
}

func loadIcon(s string) image.Image {
	f, err := os.Open(s)
	if err != nil {
		return nil
	}
	img, _, _ := image.Decode(f)
	return img
}
