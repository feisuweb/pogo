```toml
title = "Git"
date = "2016-02-04 15:00:00"
slug = "zh/docs/deploy/git"
hover = "docs"
lang = "zh"
template = "docs.html"
```

`PoGo` 可以复制本地编译内容到 git 项目并提交修改。

```bash
pogo deploy git --local="dest" --repo="repo" --message="commit {time}" --branch="master"
```

`--local` 设置本地编译好的内容的文件夹。

`--repo` 设置本地的 git 项目目录。

`--message` 设置提交的说明， `{time}` 可以标记当前的提交时间。

`--branch` 设置要提交到的 git 分支。

#### 使用步骤

- 克隆 git 项目。 `PoGo` 仅仅运行 `git push` ，不能输入用户名或密码。 因此建议克隆项目时使用 `git://` 或 `user:password` 方式。

```bash
    git clone git://your-repo.git
    // 或
    git clone https://user:password@your-repo.git
```
    
- 确认当前的分支是你需要提交的，例如 Github 一般会用 `gh-pages` 分支。

```bash
    git checkout gh-pages
```
    
- 运行 `build` 命令编译内容到 [dest] 目录。

```bash
    pogo build --dest="[dest]"
```

- 运行 `deploy` 命令提交 git 项目。

```bash
    pogo deploy --local="dest" --repo="your-repo" --branch="gh-pages"
```
    
#### 注意

`PoGo` 通过 git 的部署机制是，拷贝编译好的文件到 git 项目目录。重名的文件会被替换，其他文件不会受影响。因此有事可能需要关注 git 检测出的文件差异，保证是你需要的内容修改。