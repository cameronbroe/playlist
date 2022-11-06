package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

type Database struct {
	db *sql.DB
}

func InitializeDatabase() *Database {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./database.sqlite"
	}
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Database{
		db,
	}
}

func (db *Database) EnsureDatabaseExists() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS played_songs(
		id INTEGER PRIMARY KEY ASC,
		song_name TEXT,
		artist_name TEXT,
		album_name TEXT,
		apple_url TEXT,
		spotify_url TEXT,
		youtube_url TEXT
	);
	`

	_, err := db.db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("error ensuring database exists: %s\n", err)
		return err
	}
	return nil
}

func (db *Database) GetListOfPlayedSongs() ([]PlayedSong, error) {
	playedSongsQuery := `SELECT * FROM played_songs;`

	rows, err := db.db.Query(playedSongsQuery)
	log.Printf("got rows: %+v\n", rows)
	if err != nil {
		log.Fatalf("error querying played songs: %s\n", err)
		return []PlayedSong{}, err
	}
	log.Println("no error running query")
	log.Printf("%+v\n", rows.Err())

	playedSongs := []PlayedSong{}
	for rows.Next() {
		log.Println("foo")
		playedSong := new(PlayedSong)
		err := rows.Scan(
			&playedSong.Id, 
			&playedSong.Title, 
			&playedSong.Artist, 
			&playedSong.Album,
			&playedSong.AppleUrl, 
			&playedSong.SpotifyUrl, 
			&playedSong.YouTubeUrl,
		)
		if err != nil {
			return []PlayedSong{}, err
		}
		log.Printf("song in list: %+v\n", playedSong)
		playedSongs = append(playedSongs, *playedSong)
	}
	return playedSongs, nil
}

func (db *Database) SubmitPlayedSong(song PlayedSong) error {
	insertPlayedSongQuery := fmt.Sprintf(`
	INSERT INTO played_songs (
		song_name, artist_name, album_name, apple_url, spotify_url, youtube_url
	) VALUES (
		"%s", "%s", "%s", "%s", "%s", "%s"
	)`, song.Title, song.Artist, song.Album, song.AppleUrl, song.SpotifyUrl, song.YouTubeUrl)

	_, err := db.db.Exec(insertPlayedSongQuery)
	if err != nil {
		return err
	}
	
	return nil
}
