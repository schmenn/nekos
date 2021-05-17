package nekos_test

import (
	"fmt"
	"testing"

	"github.com/Schmenn/nekos"
	"github.com/Schmenn/nekos/nsfw"
	"github.com/Schmenn/nekos/sfw"
)

func TestMain(m *testing.M) {
	nsfwurl, err := nekos.Image(nsfw.NekoGIF)
	if err != nil {
		panic(err)
	}
	fmt.Println("NSFW URL:\t" + nsfwurl)
	sfwurl, err := nekos.Image(sfw.NekoGIF)
	if err != nil {
		panic(err)
	}
	fmt.Println("SFW URL:\t" + sfwurl)
	owotext, err := nekos.OwOify("this is a test: lorem ipsum dolor sit amet.")
	if err != nil {
		panic(err)
	}
	fmt.Println("OWO TEXT:\t" + owotext)
	spoilertext, err := nekos.Spoiler("amogus")
	if err != nil {
		panic(err)
	}
	fmt.Println("SPOILER TEXT:\t" + spoilertext)
	cat, err := nekos.Cat()
	if err != nil {
		panic(err)
	}
	fmt.Println("CAT:\t\t" + cat)
	eightball, err := nekos.EightBall()
	if err != nil {
		panic(err)
	}
	fmt.Printf("8BALL:\t\tRESPONSE:\t%s\n\t\tIMAGE URL:\t%s", eightball.Response, eightball.ImageURL)
}
