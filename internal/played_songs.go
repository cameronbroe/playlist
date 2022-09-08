package internal

type PlayedSong struct {
  Id int64 `json:"id"`
  Title string `json:"song_name"`
  Artist string `json:"artist_name"`
  Album string `json:"album_name"`
  AppleUrl string `json:"apple_url"`
  SpotifyUrl string `json:"spoitfy_url"`
  YouTubeUrl string `json:"youtube_url"`
}
