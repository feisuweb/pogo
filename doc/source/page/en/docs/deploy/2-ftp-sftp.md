```toml
title = "FTP and SFTP"
date = "2016-02-04 15:00:00"
slug = "en/docs/deploy/ftp-sftp"
hover = "docs"
lang = "en"
template = "docs.html"
```

`PoGo` can deploy static files via FTP and SFTP account. It only supports to connect remote directory with **username** and **password**.

```bash
pogo deploy ftp --local="dest" --user="user" --password="xxx" --host="127.0.0.1:21" --directory="pogo"
pogo deploy sftp --local="dest" --user="user" --password="xxx" --host="127.0.0.1:22" --directory="pogo"
```

`--local` set local directory that saving static files after building.

`--user` and `--password` set username and password for authorization to connect. In SFTP case, it doesn't support connecting via `.ssh/keys`.

`--host` set remote host address and port.

`--directory` set remote directory. In SFTP case, use `~/` as alias of user's home directory.