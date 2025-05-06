1、git 从一个仓库拉取，提交到另一个仓库

`git remote -v`

```shell
(base) ➜ ops-data-api git:(fix/bug-lfs) ✗ git remote -v
origin git@git.n.xiaomi.com:mit/operation/ops-data-api.git (fetch:拉取)
origin git@git.n.xiaomi.com:mit/operation/ops-data-api.git (push：提交)
```

可以看出，git可以单独指定两种仓库

```bash
git remote set-url origin https://gitlab.com/team/repo.git
git remote set-url --push origin https://gitlab.com/yourname/repo.git
```


