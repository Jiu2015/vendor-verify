# vendor-verify

**It is a simple tool for go mod and vendor verification.**

The go.mod file will be changed by some reasons, you need run `go mod vendor` to copy the references to vendor folder. If you forgot run `go mod vendor` then it may case out of sync, the references will backforward with go.mod, you know, it could raise an online accident...

#### How to works

In the go.mod file record the package names and versions, and the vendor root path will have modules.txt which still store the package names and versions. This tool will check both of them.

#### How to install

```shell
go get -u github.com/Jiu2015/vendor-verify/vendor-verify
```

#### How to use

```shell
vendor-verify verify --path <the path of your project>
```





