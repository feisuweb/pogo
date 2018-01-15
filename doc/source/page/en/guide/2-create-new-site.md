```toml
title = "Create New Site"
date = "2016-02-05 15:00:00"
slug = "en/guide/create-new-site"
hover = "guide"
lang = "en"
template = "guide.html"
sort = 2
```

`PoGo` packs all static assets into go source code. It can create a new site without any downloads. Run command:

```bash
pogo new site
```

Then `PoGo` extracts static files in current directory. Default site directories struct is:

    - meta.toml // save meta info for site
    - source // save all contents
    --|-- post // save all posts
    --|-- page // save all pages
    --|-- lang // language files if you need international support
    --|-- media // save media files, such as images, attachments
    - theme // save theme templates
    --|-- default // theme 'default'
    ------|-- *.html // main template files
    ------|-- embed // save embedded html template, as component parts
    ------|-- static // save static assets such as scripts, styles and images
    - lang // language files
    - media //  media files
    

## Meta Data

The basic site data are saved in `meta.toml` by default. It contains several parts to describe things to all site. Write correct data to describe your site at first.

### Meta

```toml
[meta]
# the site title
title = "PoGo"

# the site subtitle
# use to build various title for each html
subtitle = "Static Site Generator"

# the site keyword
# use in <meta content="keyword">
keyword = "pogo,golang,static,site,generator"

# the site description
# use in <meta content="description">
desc = "pogo is a simple static site generator"

# domain and root use to build correct url
# root can set subdirectory as http://domain/blog
domain = "pogo.io"
root = "http://pogo.io/"
```

### Navigation

Items of navigation are under `[[nav]]` blocks as an array in order. Full item fields:

```toml
[[nav]]
# navigator to the link
link = "/"

# link title to fill text href element
title = "Home"

# i18n key if load i18n translation
i18n = "home"

# hover class to test whether is active of this navigation item
hover = "home"

# if blank is true, it forces browser to open new tab to display the linked page
blank = true
```

### Authors

You can add several authors into the site. Then pick up one of them that assign to each page or post by author's name.

```toml
[[author]]
# author'name, must be unique
name = "pogo"

# author's email, please be private as possible
email = ""

# author's link
url = "http://pogo.io"

# author's avatar, optional. If empty, generate Gravatar image by email
avatar = ""

# author's profile 
bio = "the robot of pogo, who generates all default contents."
```

The first author set as the **Owner** of the site.

### Other format

`PoGo` support `toml`, `ini` files for `Meta Data`. Read [Format](/en/doc/cnt/format.html) documentation to get more help.