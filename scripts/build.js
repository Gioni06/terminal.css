const path = require('path')
const fs = require('fs')
const mkdirp = require('mkdirp');
const autoprefixer = require('autoprefixer')({
	browsers: [
        '>1%',
        'last 4 versions',
        'Firefox ESR',
        'not ie < 9',
      ],
      flexbox: 'no-2009'
});
const postcss      = require('postcss');
const CleanCSS = require('clean-css');
    

const css = fs.readFileSync(path.resolve(__dirname, '../src/terminal.css'), 'utf8');

mkdirp(path.resolve(__dirname,'../dist/'), function (err) {
    if (err) {
		throw e
	} else {
		postcss([ autoprefixer ]).process(css, { from: path.resolve(__dirname, '../src/terminal.css'), to: path.resolve(__dirname, '../dist/terminal.css') }).then(function (result) {
			result.warnings().forEach(function (warn) {
				console.warn(warn.toString());
				process.exit(1)
			});

			const options = {  };
			const output = new CleanCSS(options).minify(result.css);
			
			// copy to docs 
			fs.writeFileSync(path.resolve(__dirname, '../docs/terminal.min.css'), output.styles , 'utf8')
			// copy to dist   
			fs.writeFileSync(path.resolve(__dirname, '../dist/terminal.min.css'), output.styles , 'utf8')
			fs.writeFileSync(path.resolve(__dirname, '../dist/terminal.css'), result.css, 'utf8')
		});
	}
});