# Nekos
A lightweight wrapper for the nekos.life API without external dependencies

## Current Features:

all SFW and NSFW image endpoints

## TODO:
- [x] NSFW endpoints
- [x] SFW endpoints
- [ ] Text endpoints
- [ ] Better errors

## Examples:

### Installing:
`go get -u github.com/Schmenn/nekos`

### getting an Image/GIF:

```go
package main

import(
  "github.com/Schmenn/nekos" // root package
  "github.com/Schmenn/nekos/sfw" // Safe-For-Work Image endpoints
  "github.com/Schmenn/nekos/nsfw" // Not-Safe-For-Work Image endpoints
)

func main() {
  // Make a new client
  c := nekos.New()

  // get a sfw image url
  sfwURL, err := c.Image(sfw.Tickle) // "https://cdn.nekos.life/tickle/tickle_012.gif"
  if err != nil {
    // handle error
  }

  // get a nsfw image url
  nsfwURL, err := c.Image(nsfw.RandomHentaiGIF) // "https://cdn.nekos.life/Random_hentai_gif/Random_hentai_gifNB_0071.gif"
  if err != nil {
    // handle error
  }
}
