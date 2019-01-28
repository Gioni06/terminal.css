const pkg = require('./package.json')

module.exports = {
  build: {
    srcPath: './src',
    outputPath: './public'
  },
  site: {
    title: 'Terminal CSS',
    libVersion: pkg.version,
    description: pkg.description,
    keywords: pkg.keywords
  }
};
