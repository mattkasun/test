function chooseBinary() {
    // ...
        return `main-linux-amd64-${VERSION}`
    // ...
}

const binary = chooseBinary()
const mainScript = `${__dirname}/${binary}`
const spawnSyncReturns = childProcess.spawnSync(mainScript, { stdio: 'inherit' })
