package nekos

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

// NSFWEndpoint NSFW Endpoint
type NSFWEndpoint string

// SFWEndpoint SFW Endpoint
type SFWEndpoint string

// Image sends a request to the nekos.life api and returns the image url for the given endpoint
func Image(endpoint interface{}) (string, error) {
	if v, ok := endpoint.(NSFWEndpoint); ok {
		return imgReq(string(v))
	} else if v, ok := endpoint.(SFWEndpoint); ok {
		return imgReq(string(v))
	}
	return "", errors.New("Wrong Endpoint provided")
}

// OwOify OwOifies text
func OwOify(text string) (string, error) {
	res, err := http.Get("https://nekos.life/api/v2/owoify?text=" + url.QueryEscape(text))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var r struct {
		Text string `json:"owo"`
	}
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

func imgReq(ep string) (string, error) {
	res, err := http.Get("https://nekos.life/api/v2" + string(ep))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var response struct {
		URL string `json:"url"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	return response.URL, nil
}
