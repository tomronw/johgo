const express = require('express');
const path = require('path');
const { createProxyMiddleware } = require('http-proxy-middleware');

const app = express();

app.use(express.static(path.join(__dirname, 'build')));

app.use('/v1/search', createProxyMiddleware({
  target: 'http://localhost:16078',
  changeOrigin: true
}));

app.use('/v1/search_suggestions', createProxyMiddleware({
  target: 'http://localhost:8005',
  changeOrigin: true
}));

app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, 'build', 'index.html'));
});

const port = process.env.PORT || 80;
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});