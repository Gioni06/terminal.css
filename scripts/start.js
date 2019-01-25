const chokidar = require('chokidar');
const liveServer = require('live-server');
const build = require('./utils/build-fn');
const path = require('path');

function debounce(func, wait, immediate) {
	var timeout;
	return function() {
		var context = this, args = arguments;
		var later = function() {
			timeout = null;
			if (!immediate) func.apply(context, args);
		};
		var callNow = immediate && !timeout;
		clearTimeout(timeout);
		timeout = setTimeout(later, wait);
		if (callNow) func.apply(context, args);
	};
};

/**
 * Serve the site in watch mode
 */
const serve = (flags) => {
  console.log(`Starting local server at http://localhost:${flags.port}`);

  const options = {
	sourceFile: path.resolve(__dirname, '../src/terminal.css'),
	distFolder: path.resolve(__dirname, '../dist'),
	docsFolder: path.resolve(__dirname, '../docs')
  }

  build.run(options);
  liveServer.start({
      port: flags.port,
      root: options.docsFolder,
      open: true,
      logLevel: 0
    });

  chokidar.watch(options.sourceFile, { ignoreInitial: true }).on(
    'all',
    debounce(() => {
      build.run(options);
      console.log('Waiting for changes...');
    }, 500)
  );
};

serve({ port: 3000 })
