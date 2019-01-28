const cpFile = require('cp-file');
const path = require('path');

(async () => {
    await cpFile(path.resolve(__dirname, '../src/browserconfig.xml'), './public/browserconfig.xml');
    await cpFile(path.resolve(__dirname, '../src/manifest.json'), './public/manifest.json');
    console.log('browserconfig and manifest copied');
})();