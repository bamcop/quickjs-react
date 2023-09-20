# README

## Why?
react islands with golang

## examples
### render_to_string
The minimal component for validating react SSR in golang + quickjs

#### 1. run
```sh
bun run index.js
```

#### 2. bundle use esbuild cmd
```sh
npx esbuild --loader:.js=jsx --bundle index.js --outfile=index.bundle_esbuild_cli.js
```

#### 3. bundle use esbuild golang api
```sh
go run examples/bundle/main.go
```
