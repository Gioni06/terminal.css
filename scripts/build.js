const build = require('./utils/build-fn')
const path = require('path')
const nanogen = require('nanogen')
const staticSiteOptions = require('../site.config');

const options = {
    sourceFile: path.resolve(__dirname, '../lib/terminal.css'),
    distFolder: path.resolve(__dirname, '../dist'),
    docsFolder: path.resolve(__dirname, '../public'),
    docsSrcFolder: path.resolve(__dirname, '../src')
  }

build.run(options);
nanogen.build({ site: staticSiteOptions.site })
