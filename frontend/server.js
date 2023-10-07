const express = require('express');
const path = require('path');
const { createProxyMiddleware } = require('http-proxy-middleware');

const app = express();
require('dotenv').config();

app.use(express.static(path.join(__dirname, 'build')));

app.use('/search', createProxyMiddleware({
  target: process.env.REACT_APP_APIENDPOINT,
  changeOrigin: true
}));

app.use('/search_suggestions', createProxyMiddleware({
  target: process.env.REACT_APP_BASEAPI,
  changeOrigin: true
}));

app.use('/serve_calendar', createProxyMiddleware({
  target: process.env.REACT_APP_BASEAPI,
  changeOrigin: true
}));

app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});

const port = process.env.PORT || 80;
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});