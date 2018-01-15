```toml
title = "Download & Compile"
date = "2016-02-04 15:00:00"
slug = "en/docs/dl-compile"
hover = "docs"
lang = "en"
template = "docs.html"
```

To install `PoGo`, you must have `Go` language. Download **Go** and install it from [official site](https://golang.org).

`PoGo` need **Go 1.4+**.

If Go is ready, get `PoGo`'s source code:

```go
go get github.com/feisuweb/pogo
```

Then you can modify source codes and build your own binary.

### Asset

`PoGo` use [go-bindata](https://github.com/jteeuwen/go-bindata) to bundle assets to source code.

```go
go get github.com/jteeuwen/go-bindata
```

Add `$GOPATH/bin` to `PATH`, then:

```go
go generate -x
```

It packs `doc`, `source`, `theme` to go source code in `app/asset/asset.go`.

### Notice

- `PoGo` need `fsnotify` to watch file changes. It's not tested on **arm** platform.
- `PoGo` use `log15.v2` to print logs. But it can't be compiled in **openbsd** & **netbsd**.