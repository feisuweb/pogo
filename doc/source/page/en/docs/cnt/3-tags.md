```toml
title = "Tags"
date = "2016-02-04 15:00:00"
slug = "en/docs/cnt/tags"
hover = "docs"
lang = "en"
template = "docs.html"
```

Post provides tags to classify. In post front-matter:

```toml
tags = ["pogo","golang"]
```

Or ini format:

```ini
tags = "pogo,golang"
```

`PoGo` generates html files for each tag with list of tagged posts order by time but not paged.

