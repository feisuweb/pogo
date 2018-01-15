```toml
title = "About PoGo"
# slug = "about"
desc = "some words about pogo"
date = "2016-03-24 12:24:00"
author = "pogo"
# set nav to active status when this page
hover = "about"
# set template file to render this page
template = ""

[meta]
metadata = "this is meta data"
```

### Introduction

`PoGo` is a simple static site generator by [Golang](https://golang.org). It compiles [markdown](https://help.github.com/articles/markdown-basics/) to site pages with beautiful theme. No dependencies, cross platform and very fast.

![golang](@media/golang.png)

### Quick start

1. Download from [PoGo](http://pogo.io) and extract zip archive.
2. Run `pogo new site` to create new default site.
2. Run `pogo server`, open `http://localhost:9899` to visit.


### Commands

Run a command when run `pogo` executable file:

- `pogo new` create new site, post or page.
- `pogo build` build static files.
- `pogo server` build and serve static files.

More details in [Documentation](http://pogo.io/en/docs.html).

### Writing

`PoGo` support two kinds of content, `post` and `page`. you can create any `.md` file in proper directories in `source` directory. Read the [wiki](http://pogo.io/en/guide/write-new-post.html) to learn the layout format and more details.

### Compile

After you change your source `.md` files, run

    pogo build

To build static files.

### Customize

- Theme: theme documentation is [Here](http://pogo.io/en/docs/tpl/syntax.html)
