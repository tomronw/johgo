// App.js
import React from 'react';
import { SearchProvider } from './context/SearchContext';
import { useTheme } from './context/ThemeContext';
import { Routes, Route } from "react-router-dom";
import Mainbar from './components/toolbar/Mainbar';
import Home from './pages/Home';
import NotFound from './pages/NotFound';
import './css/Theme.css';
import Footer from './components/miscellaneous/Footer';
import { SuggestionProvider } from './context/SuggestionsContext';
function App() {
  const { darkMode } = useTheme();

  return (
    <SearchProvider>
      <SuggestionProvider>
        <div className={`App ${darkMode ? 'dark-mode' : 'light-mode'}`}>
          <Mainbar />
          <Routes>
            <Route exact path="/" element={<Home />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </div>
        <Footer />
      </SuggestionProvider>
    </SearchProvider>
  );
}

export default App;

