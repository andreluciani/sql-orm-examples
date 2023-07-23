const fs = require('fs');
const cheerio = require('cheerio');

const scriptToInject = `<script async src="https://umami-andre-luciani.up.railway.app/script.js" data-website-id="${process.env.UMAMI_WEBSITE_ID}"></script>`;

const indexPath = './dist/index.html';

fs.readFile(indexPath, 'utf8', (err, data) => {
    if (err) {
        console.error('Error reading index.html:', err);
        return;
    }

    const $ = cheerio.load(data);
    $('head').append(scriptToInject);

    fs.writeFile(indexPath, $.html(), 'utf8', (writeErr) => {
        if (writeErr) {
            console.error('Error writing index.html:', writeErr);
            return;
        }

        console.log('Umami analytics added successfully into index.html');
    });
});
