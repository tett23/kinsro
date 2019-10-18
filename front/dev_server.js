const express = require('express');

const app = express();
const port = 4355;

app.use(express.static('./public'));
app.get('/vindex', (_, res) => {
  res.sendFile(process.env['VINDEX_PATH'], {
    headers: {
      'Content-Type': 'application/octet-stream',
    },
  });
});

app.listen(port, () => console.log(`Example app listening on port ${port}!`));
