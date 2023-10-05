function chooseBinary() {
    // ...
        return `main-linux-amd64-v1.0.0`
    // ...
}
import core from '@actions/core';
import child_process from '@node/child_process'
try {
const binary = chooseBinary();
const mainScript = `${__dirname}/${binary}`;
const args = getInput('cmd-args');
var spawnSyncReturns = child_process.spawnSync(mainScript, args, { stdio: 'inherit' });
console.log('stdout:\n'+spawnSyncReturns.stdout);
} catch (error) {
    core.setFailed(error.message);
}
