package nekos_test

import (
	"fmt"
	"testing"

	"github.com/Schmenn/nekos"
	"github.com/Schmenn/nekos/nsfw"
	"github.com/Schmenn/nekos/sfw"
)

func TestMain(m *testing.M) {
	c := nekos.New()
	nsfwurl, err := c.Image(nsfw.NekoGIF)
	if err != nil {
		panic(err)
	}
	fmt.Println(nsfwurl)
	sfwurl, err := c.Image(sfw.NekoGIF)
	if err != nil {
		panic(err)
	}
	fmt.Println(sfwurl)
}
