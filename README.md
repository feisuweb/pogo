# Pogo 破狗
一款简单快速的静态网站生成工具，采用golang编写，第一个版本是付小黑的pugo.io,
本版本进行了大幅度的改版与更新，故单独取名叫做pogo，中文名称叫破狗。

# 如何编译


# 如果修改界面

由于界面采用了集成到程序中的方式，所以在每次修改模板的时候，需要进行一次UI代码的生成。
命令如下：
```
go-bindata -o=app/asset/asset.go -pkg=asset source/... doc/source/... doc/theme/...
```

将代码重新格式化

```
gofmt -w -s .
```
# 如何使用
## 创建一个静态站点
名称叫做 xxx.com，同步创建一个xxx.com的目录
```
pogo new site xxx.com
```
## 实时监控，在xxx.com目录下，执行如下命令
```
pogo server
```
打开浏览器，输入http://localhost:9899/
## 通过WEB进行文章管理
打开浏览器，输入http://localhost:9899/pogoadmin/
## 创建一篇文章
```
pogo new post 文章英文名称
```
## 创建一个页面
```
pogo new page 页面英文名称
```
## 发布到FTP服务器上
举个例子，将本地的站点发布到服务上，服务器的路径是`/domains/www.feisuweb.com/public_html/`

```
pogo deploy ftp --user FTP用户名 --password FTP密码 --host www.feisuweb.com:21 --directory /domains/www.feisuweb.com/public_html/ --debug
```


