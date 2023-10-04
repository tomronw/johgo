import React from 'react';
import ReactDOM from 'react-dom/client';
import './css/Index.css';
import App from './App';
import { ThemeProvider } from './context/ThemeContext'; // Update the path
import { BrowserRouter } from "react-router-dom";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <ThemeProvider>
        <App />
      </ThemeProvider>
    </BrowserRouter>
  </React.StrictMode>
);