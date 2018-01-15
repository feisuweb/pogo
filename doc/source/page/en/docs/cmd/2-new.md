```toml
title = "New"
date = "2016-02-04 15:00:00"
slug = "en/docs/cmd/new"
hover = "docs"
lang = "en"
template = "docs.html"
```

`new` command create new `site`, `post` or `page`.

### Site

Assets are bundled in `PoGo`. Just it can create new site without any downloads:

```go
pogo new site
```

It extracts common defaults, starting posts and pages in `source` and three themes in `theme` directory in current directory. 

```go
pogo new site --doc
```

When set `--doc` flag, it extracts `doc` directory that contains all documentation data. The data is ready to compile to document website. Read [Doc](/en/docs/cmd/doc) command to get more help.

### Post

Create new post without title:

```go
pogo new post
```

Default post markdown file is created in `source/[year]/[day-month-hour-minute-second].md`

Try with post title:

```go
pogo new post "this is new post"
```

Now the file is `source/this-is-new-post.md`.

### Page

Same usage to `new post`:

```go
pogo new page
pogo new page "this is new page"
```