import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import './App.css';
import { Playlist } from './Playlist';

const queryClient = new QueryClient();

function App() {
  return (
    <div className="app">
      <header>
        <nav>
          <span><a href="https://cameronbroe.com">[ my homepage ]</a></span>
          <span><a href="https://github.com/cameronbroe/playlist">[ github ]</a></span>
        </nav>
        <div>playlist.cameronbroe.com</div>
      </header>
      <QueryClientProvider client={queryClient}>
        <Playlist />
      </QueryClientProvider>
      <footer>
        <div>Made with 💙 in Memphis, TN</div>
        <div>Color scheme based on <a href="https://ethanschoonover.com/solarized/">solarized</a></div>
        <div>Streaming service data filled in from <a href="https://odesli.co">Odesli</a></div>
      </footer>
    </div>
  );
}

export default App;
