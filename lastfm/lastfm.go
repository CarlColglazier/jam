// Copyright (c) 2016, 2017 Evgeny Badin

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package lastfm

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/shkh/lastfm-go/lastfm"
)

// Client provides a Last.FM API client.
type Client struct {
	Api      *lastfm.Api
	Username string
	Password string
}

func New(apiKey string, apiSecret string) *Client {
	return &Client{
		Api: lastfm.New(apiKey, apiSecret),
	}
}

func (client *Client) GetToken() (string, error) {
	token, err := client.Api.GetToken()
	if err != nil {
		return "", err
	}
	authURL := client.Api.GetAuthTokenUrl(token)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("If you use Last.fm click the link below, grant permission to Jam, and press enter; if not, just press enter:\n", authURL)
	_, _ := reader.ReadString('\n')
	return token, nil
}

func (client *Client) LoginWithToken(token string) error {
	return client.Api.LoginWithToken(token)
}
func (client *Client) Scrobble(artist, track string, timestamp int64) error {
	p := lastfm.P{"artist": artist, "track": track,
		"timestamp": timestamp}
	_, err := client.Api.Track.Scrobble(p)
	if err != nil {
		log.Printf("Error scrobble: %s", err)
		return err
	}
	return nil
}

func (client *Client) NowPlaying(track, artist string) error {
	p := lastfm.P{"artist": artist, "track": track}
	_, err := client.Api.Track.UpdateNowPlaying(p)
	if err != nil {
		log.Printf("Error submitting now playing: %s", err)
		return err
	}
	return nil
}
