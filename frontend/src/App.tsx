import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import './App.css';
import { Playlist } from './Playlist';

const queryClient = new QueryClient();

function App() {
  return (
    <div className="app">
      <header>playlist.cameronbroe.com</header>
      <QueryClientProvider client={queryClient}>
        <Playlist />
      </QueryClientProvider>
    </div>
  );
}

export default App;
