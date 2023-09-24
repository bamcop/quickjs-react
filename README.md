# README
Run React SSR with Golang using QuickJS

## Run
```sh
go run cmd/dev/dev.go
```

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

### summary
```shell
(base) ➜ go run examples/qucikjs_render/main.go
<h1>Hello, world!</h1><button>You clicked me <!-- -->0<!-- --> times</button>
(base) ➜ node examples/render_to_string/index.bundle_esbuild_cli.js
<h1>Hello, world!</h1><button>You clicked me <!-- -->0<!-- --> times</button>
(base) ➜ node examples/render_to_string/index.bundle_esbuild_go_api.js
<h1>Hello, world!</h1><button>You clicked me <!-- -->0<!-- --> times</button>
```

### Benchmark
```shell
hyperfine --warmup 5 --runs 25 'node index.bundle_esbuild_cli.js'
Benchmark 1: node index.bundle_esbuild_cli.js
  Time (mean ± σ):      35.5 ms ±   0.2 ms    [User: 31.2 ms, System: 4.0 ms]
  Range (min … max):    35.0 ms …  36.0 ms    25 runs
```

```shell
go test -bench BenchmarkEvalRenderToString .
goos: darwin
goarch: arm64
pkg: github.com/bamcop/quickjs-react/examples/qucikjs_render
BenchmarkEvalRenderToString-8
      57	  20126838 ns/op
PASS
ok  	github.com/bamcop/quickjs-react/examples/qucikjs_render	2.323s
```
