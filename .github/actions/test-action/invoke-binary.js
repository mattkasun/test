function chooseBinary() {
    // ...
        return `main-linux-amd64-v1.0.0`
    // ...
}
const  child_process   = require('node:child_process');
const core = require('@actions/core')
try {
const binary = chooseBinary();
const mainScript = `${__dirname}/${binary}`;
console.log("bin", mainScript)
const args = core.getInput('cmd-args');
console.log("args", args)
var spawnSyncReturns = child_process.spawnSync(mainScript, [args], { stdio: 'inherit' });
console.log('stdout:\n'+spawnSyncReturns.stdout);
} catch (error) {
    core.setFailed(error.message);
}
