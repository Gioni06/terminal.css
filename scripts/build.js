const build = require('./utils/build-fn')
const path = require('path')

build.run({
	sourceFile: path.resolve(__dirname, '../src/terminal.css'),
	distFolder: path.resolve(__dirname, '../dist'),
	docsFolder: path.resolve(__dirname, '../docs'),
})