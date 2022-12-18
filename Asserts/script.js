const golang = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), golang.importObject)
.then((result) => {
    golang.run(result.instance)
})