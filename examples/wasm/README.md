# README
[Wasm By Example](https://wasmbyexample.dev/examples/hello-world/hello-world.go.en-us.html)

```sh
tinygo build -o main.wasm -target wasm ./main.go
```

```sh
cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js .
```

```sh
http-server -p 3009
```