package nekos

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
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
	}
	return "", errors.New("Wrong Endpoint provided")
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
	body, err := io.ReadAll(res.Body)
	var response urlresponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	return response.URL, nil
}

func (c *Client) OwOify(text string) (string, error) {
	res, err := http.Get("https://nekos.life/api/v2/owoify?text=" + url.QueryEscape(text))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	r := struct {
		Text string `json:"owo"`
	}{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}
	return r.Text, nil
}
