```toml
title = "Standalone"
date = "2016-02-04 15:00:00"
slug = "en/docs/deploy/standalone"
hover = "docs"
lang = "en"
template = "docs.html"
```

`PoGo` can run http server by self. In production, use `server` command to setup static files after building:

```bash
pogo build
pogo server --static
```

If in writing mode, use `server` directory to watch changes and update immediately.

```bash
pogo server
```