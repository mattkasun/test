function chooseBinary() {
    // ...
        return `main-linux-amd64-v1.0.0`
    // ...
}
const core = require('@actions/core');
try {
var  child_process   = require('node:child_process');
const binary = chooseBinary();
const mainScript = `${__dirname}/${binary}`;
const args = core.getInput('cmd-args');
var spawnSyncReturns = child_process.spawnSync(mainScript, args, { stdio: 'inherit' });
console.log('stdout:\n'+spawnSyncReturns.stdout);
} catch (error) {
    core.SetFailed(error.message);
    
}
