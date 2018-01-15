```toml
title = "Build Files"
date = "2016-02-05 15:00:00"
slug = "en/guide/build-files"
hover = "guide"
lang = "en"
template = "guide.html"
sort = 6
```

Use `build` command to build files:

```bash
pogo build 
```

### Watch

`PoGo` can watch changes and re-build files immediately. It overwrites any html files and checks md5sum to replace static files that needed.

```bash
pogo build --watch
```

### Custom Source

Build files from custom source:

```bash
pogo build --source="your-source"
```

### Custom Destination

Build files to custom destination:

```bash
pogo build --dest="your-directory"
```

### Custom Theme

Build files with specific theme:

```bash
pogo build --theme="your-theme-directory"
```