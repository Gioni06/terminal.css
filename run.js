const { spawn } = require('child_process');
const os = require('os');

// Define the base name of your application
const appName = 'terminalcss-builder';

// Get the OS and architecture
const platform = os.platform(); // 'darwin', 'linux', 'win32', etc.
const arch = os.arch(); // 'x64', 'arm64', 'ia32' (for x86), etc.

// Map Node.js architecture and platform strings to your binary naming convention
const platformMap = {
    'win32': 'windows',
    'darwin': 'darwin',
    'linux': 'linux'
};
const archMap = {
    'x64': 'amd64',
    'arm64': 'arm64',
    'ia32': '386' // ia32 represents x86 in Node.js
};

// Construct the binary name based on the platform and architecture
let binaryName = `./builds/${appName}-${platformMap[platform]}-${archMap[arch]}`;
if (platform === 'win32') {
    binaryName += '.exe';
}

// Add command line arguments to the binary command
const args = process.argv.slice(2).join(' '); // Skip the first two elements

// Spawn the binary process
const binaryProcess = spawn(binaryName, [args]);

// Forward stdout and stderr of the binary to the Node.js process
binaryProcess.stdout.on('data', (data) => {
    process.stdout.write(data);
});

binaryProcess.stderr.on('data', (data) => {
    process.stderr.write(data);
});

// Handle binary process exit
binaryProcess.on('close', (code) => {
    console.log(`Binary process exited with code ${code || 0}.`);
    process.exit(code);
});

// Function to terminate the binary process gracefully
const terminateBinaryProcess = () => {
    console.log('Terminating binary process...');
    binaryProcess.kill(); // Sends SIGTERM to the binary process
};

// Listen for termination signals
process.on('SIGINT', terminateBinaryProcess); // Ctrl+C signal
process.on('SIGTERM', terminateBinaryProcess); // Termination signal
