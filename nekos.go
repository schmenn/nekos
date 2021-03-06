package nekos

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// NSFWEndpoint NSFW Endpoint
type NSFWEndpoint string

// SFWEndpoint SFW Endpoint
type SFWEndpoint string

// Client nekos.go client
type Client struct {
	httpClient http.Client
}

type urlresponse struct {
	// URL URL of the image
	URL string `json:"url"`
}

// New creates a client
func New() (c Client) {
	c = Client{httpClient: http.Client{Timeout: time.Second * 5}}
	return
}

// Image sends a request to the nekos.life api and returns the image url for the given endpoint
func (c *Client) Image(endpoint interface{}) (string, error) {

	if v, ok := endpoint.(NSFWEndpoint); ok {
		return imgReq(c, string(v))
	} else if v, ok := endpoint.(SFWEndpoint); ok {
		return imgReq(c, string(v))
	} else {
		log.Fatalln("You provided the wrong endpoint.")
	}
	return "", nil
}

func imgReq(c *Client, ep string) (string, error) {
	req, err := http.NewRequest("GET", "https://nekos.life/api/v2"+string(ep), nil)
	if err != nil {
		return "", err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var response urlresponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	return response.URL, nil
}
