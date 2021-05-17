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

// Spoiler turns text into a spoiler-packed text
func Spoiler(text string) (string, error) {
	res, err := http.Get("https://nekos.life/api/v2/spoiler?text=" + url.QueryEscape(text))
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

// Why returns random question
func Why() (string, error) {
	var j struct {
		Text string `json:"why"`
	}
	res, err := http.Get("https://nekos.life/api/v2/why")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(body, &j)
	return j.Text, nil
}

// Cat returns random cat
func Cat() (string, error) {
	var j struct {
		Text string `json:"cat"`
	}
	res, err := http.Get("https://nekos.life/api/v2/cat")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(body, &j)
	return j.Text, nil
}

// EightBall returns random response and image of an 8ball
func EightBall() (struct{ Response, ImageURL string }, error) {
	var j struct {
		Response string `json:"response"`
		ImageURL string `json:"url"`
	}
	res, err := http.Get("https://nekos.life/api/v2/8ball")
	if err != nil {
		return struct {
			Response string
			ImageURL string
		}{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return struct {
			Response string
			ImageURL string
		}{}, err
	}
	json.Unmarshal(body, &j)
	return struct {
		Response string
		ImageURL string
	}(j), nil
}

// Fact returns random fact
func Fact() (string, error) {
	var j struct {
		Text string `json:"fact"`
	}
	res, err := http.Get("https://nekos.life/api/v2/fact")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(body, &j)
	return j.Text, nil
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
