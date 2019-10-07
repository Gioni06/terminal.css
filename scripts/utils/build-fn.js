const path = require('path')
const fs = require('fs')
const mkdirp = require('mkdirp');

function run({
	sourceFile,
	distFolder,
	docsFolder
}) {
	const autoprefixer = require('autoprefixer')({});
	const postcss      = require('postcss');
	const CleanCSS = require('clean-css');
		

	const css = fs.readFileSync(sourceFile, 'utf8');

	mkdirp(distFolder, function (err) {
		if (err) {
			throw e
		} else {
			postcss([ autoprefixer ]).process(css, { from: sourceFile, to: path.resolve(distFolder, 'terminal.css') }).then(function (result) {
				result.warnings().forEach(function (warn) {
					console.warn(warn.toString());
					process.exit(1)
				});

				const options = {  };
				const output = new CleanCSS(options).minify(result.css);
				
				// copy to docs 
				fs.writeFileSync(path.resolve(docsFolder, 'terminal.min.css'), output.styles , 'utf8')
				// copy to dist   
				fs.writeFileSync(path.resolve(distFolder, 'terminal.min.css'), output.styles , 'utf8')
				fs.writeFileSync(path.resolve(distFolder, 'terminal.css'), result.css, 'utf8')
			});
		}
	});
}

module.exports = {
	run
}