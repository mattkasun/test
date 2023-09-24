function chooseBinary() {
    // ...
        return `main-linux-amd64-v1.0.0`
    // ...
}

const binary = chooseBinary()
const mainScript = `${__dirname}/${binary}`
const spawnSyncReturns = childProcess.spawnSync(mainScript, { stdio: 'inherit' })
