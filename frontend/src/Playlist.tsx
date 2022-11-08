import './Playlist.css';
import { useQuery } from "@tanstack/react-query";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSpotify, faYoutube, faApple } from '@fortawesome/free-brands-svg-icons'

interface PlaylistItem {
  song_name: string;
  artist_name: string;
  album_name: string;
  apple_url: string;
  spotify_url: string;
  youtube_url: string;
  album_art_url: string;
  id: number;
}

export function Playlist() {
  console.log(process.env.REACT_APP_PLAYLIST_LIST_URL);
  const { isLoading, error, data, isSuccess } = useQuery<PlaylistItem[]>({
    queryKey: ['playlistData'],
    queryFn: async () => {
      return fetch(process.env.REACT_APP_PLAYLIST_LIST_URL!).then(res => {
        return res.json();
      });
    }
  });

  if(isLoading) return <div>Playlist is loading</div>;

  if(error && error instanceof Error) return <div>Error has occurred while loading playlist: {JSON.stringify(error)}</div>;

  return (
    <div className="playlist-container">
      {isSuccess && data.map(item =>
        <div className="playlist-item" key={item.id}>
          <div className="art"><img src={item.album_art_url} alt="" /></div>
          <div className="playlist-item-details">
            <div className="title">{item.song_name}</div>
            <div className="artist">{item.artist_name}</div>
            <div className="album">{item.album_name}</div>
          </div>
          <div className="playlist-item-links">
            <div><a href={item.spotify_url}><FontAwesomeIcon icon={faSpotify} /></a></div>
            <div><a href={item.apple_url}><FontAwesomeIcon icon={faApple} /></a></div>
            <div><a href={item.youtube_url}><FontAwesomeIcon icon={faYoutube} /></a></div>
          </div>
        </div>
      )}
    </div>
  );
}
