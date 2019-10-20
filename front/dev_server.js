const express = require('express');
const path = require('path');

const app = express();
const port = 4355;

app.use(express.static('./public'));
app.get('/contents/*', (_, res) => {
  res.sendFile(path.join(__dirname, 'public/index.html'));
});
app.get('/vindex', (_, res) => {
  res.sendFile(process.env['VINDEX_PATH'], {
    headers: {
      'Content-Type': 'application/octet-stream',
    },
  });
});

app.listen(port, () => console.log(`Example app listening on port ${port}!`));
