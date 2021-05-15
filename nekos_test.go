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
	fmt.Println(nsfwurl)
	sfwurl, err := nekos.Image(sfw.NekoGIF)
	if err != nil {
		panic(err)
	}
	fmt.Println(sfwurl)
	owotext, err := nekos.OwOify("this is a test: lorem ipsum dolor sit amet.")
	if err != nil {
		panic(err)
	}
	fmt.Println(owotext)
}
