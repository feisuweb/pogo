```toml
title = "Build"
date = "2016-02-04 15:00:00"
slug = "en/docs/cmd/build"
hover = "docs"
lang = "en"
template = "docs.html"
```

`build` command can build all contents to website html files。

### Just Build

`build` command basic usage:

```go
pogo build [--source="source"] [--dest="dest"] [--theme="theme/default"] [--watch] [--debug]
```

`--source` set the source directory, default is `source`.

`--dest` set the directory that PoGo builds contents to, default is `dest`.

`--theme` set the directory of theme, default is `theme/default` ( PoGo provides 3 themes in `theme` ).

`--watch` set flag to watching changes and rebuild site.

`--debug` print more logs when running command.

