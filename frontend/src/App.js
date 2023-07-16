import './css/App.css';
import MainAppBar from './components/MainAppBar.jsx';
import { Routes, Route } from "react-router-dom";
import About from './pages/About';
import Search from './pages/Search';
import Footer from './components/Footer';
import { SearchProvider } from './context/SearchContext'
import NotFound from './pages/NotFound';
import { useState } from 'react';

function App() {

  const [isLightMode, setIsLightMode] = useState(false);

  const handleLightModeChange = (newMode) => {
    setIsLightMode(newMode);
  };


  return (
    <SearchProvider>
      <div className={isLightMode ? 'dark-bg' : 'light-bg'}>
        <div className="App">
          <MainAppBar setLightMode={handleLightModeChange} />
          <Routes>
            <Route exact path="/" element={<Search />}>
            </Route>
            <Route>
              <Route path="/about" element={<About />}>
              </Route>
            </Route>
            <Route>
              <Route path="*" element={<NotFound />} />
            </Route>
          </Routes>
        </div>
        <Footer />
      </div>
    </SearchProvider>
  );
}

export default App;
