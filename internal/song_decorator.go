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
	ArtistName string `json:"artistName"`
	AlbumName  string `json:"collectionName"`
	TrackName  string `json:"trackName"`
	TrackUrl   string `json:"trackViewUrl"`
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

	for _, result := range(results.Results) {
		artistMatches := normalize(result.ArtistName) == normalize(song.Artist)
		albumMatches := normalize(result.AlbumName) == normalize(song.Album)
		titleMatches := normalize(result.TrackName) == normalize(song.Title)

		if artistMatches && titleMatches && albumMatches {
			return &result, nil
		}
	}
	return nil, errors.New("could not find track on iTunes")
}

func (sd *SongDecorator) DecoratePlayedSong(song *PlayedSong) {
	itunesResult, err := searchItunes(song)
	if err != nil {
		fmt.Printf("could not find song on iTunes: %s\n", err)
		return
	}

	song.AppleUrl = itunesResult.TrackUrl
}
