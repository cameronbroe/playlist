package internal

import (
	"database/sql"
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
    dbPath = "./database.db"
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
	playedSongsQuery := `
	SELECT 
		id, song_name, artist_name, album_name, apple_url, spotify_url, youtube_url
	FROM played_songs;
	`
	
	rows, err := db.db.Query(playedSongsQuery)
	if err != nil {
		log.Fatalf("error querying played songs: %s\n", err)
		return []PlayedSong{}, err
	}

	playedSongs := []PlayedSong{}
	for rows.Next() {
		playedSong := new(PlayedSong)
		err := rows.Scan(playedSong)
		if err != nil {
			return []PlayedSong{}, err
		}
		playedSongs = append(playedSongs, *playedSong)
	}
	return playedSongs, nil
}
