```toml
title = "New"
date = "2016-02-04 15:00:00"
slug = "zh/docs/cmd/new"
hover = "docs"
lang = "zh"
template = "docs.html"
```

`new` 命令可以创建新的 `site`( 站点 ), `post`( 文章 ) 或 `page`( 页面 )。

### 站点

`PoGo` 内嵌有静态资源，因而可以直接新建站点，不需要下载任何附加内容：

```go
pogo new site
```

释放好的新站点内有默认的配置，`source` 文件夹有起始文章和页面，`theme` 文件夹有三个主题。

```go
pogo new site --doc
```

当设置 `--doc`，将只释放 `doc` 目录带有所有的文档资源。你可以使用 [Doc](/zh/docs/cmd/doc) 命令浏览文档内容。


### 文章

直接创建文章：

```go
pogo new post
```

空文章会生成在 `source/[year]/[day-month-hour-minute-second].md` 文件。

创建带标题的文章：

```go
pogo new post "this is new post"
```

新文章保存在 `source/this-is-new-post.md` 文件。

### 页面

和 `new post` 操作一样：

```go
pogo new page
pogo new page "this is new page"
```