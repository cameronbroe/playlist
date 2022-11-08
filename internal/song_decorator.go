package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SongDecorator struct{}

type itunesResult struct {
	ArtistName  string `json:"artistName"`
	AlbumName   string `json:"collectionName"`
	TrackName   string `json:"trackName"`
	TrackUrl    string `json:"trackViewUrl"`
	AlbumArtUrl string `json:"artworkUrl100"`
}

type itunesResponse struct {
	Results []itunesResult `json:"results"`
}

func normalize(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

func searchItunes(song *PlayedSong) (*itunesResult, error) {
	escapedQuery := url.QueryEscape(fmt.Sprintf("%s %s", song.Title, song.Artist))
	itunesUrl := fmt.Sprintf("https://itunes.apple.com/search?term=%s&entity=musicTrack", escapedQuery)
	resp, err := http.Get(itunesUrl)
	if err != nil {
		log.Printf("error searching for song on iTunes: %s\n", err)
		return nil, err
	}

	var results itunesResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&results)

	for _, result := range results.Results {
		artistMatches := normalize(result.ArtistName) == normalize(song.Artist)
		albumMatches := normalize(result.AlbumName) == normalize(song.Album)
		titleMatches := normalize(result.TrackName) == normalize(song.Title)

		if artistMatches && titleMatches && albumMatches {
			return &result, nil
		}
	}
	return nil, errors.New("could not find track on iTunes")
}

type odesliResult struct {
	OdesliUrl string `json:"pageUrl"`
	Platforms struct {
		Spotify struct {
			Url string `json:"url"`
		} `json:"spotify"`

		YouTube struct {
			Url string `json:"url"`
		} `json:"youtube"`
	} `json:"linksByPlatform"`
}

func searchOdesli(song *PlayedSong) (*odesliResult, error) {
	escapedUrl := url.QueryEscape(song.AppleUrl)
	log.Println(escapedUrl)
	odesliUrl := fmt.Sprintf("https://api.song.link/v1-alpha.1/links?url=%s", escapedUrl)
	resp, err := http.Get(odesliUrl)
	if err != nil {
		log.Printf("error getting data from Odesli: %s\n", err)
		return nil, err
	}

	var result odesliResult
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&result)
	return &result, nil
}

func (sd *SongDecorator) DecoratePlayedSong(song *PlayedSong) {
	itunesResult, err := searchItunes(song)
	if err != nil {
		fmt.Printf("could not find song on iTunes: %s\n", err)
		return
	}

	song.AppleUrl = itunesResult.TrackUrl
	song.AlbumArtUrl = itunesResult.AlbumArtUrl

	odesliResult, err := searchOdesli(song)
	if err != nil {
		fmt.Printf("could not find song on Odesli: %s\n", err)
		return
	}

	song.SpotifyUrl = odesliResult.Platforms.Spotify.Url
	song.YouTubeUrl = odesliResult.Platforms.YouTube.Url
}
