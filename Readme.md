## memo
- `go.mod`は手動ではなく、コマンドで作る
    - コマンド: `$ go mod init GoAndMySQL`
    - 結果: `go: creating new go.mod: module GoAndMySQL`
    - 参考: https://qiita.com/uchiko/items/64fb3020dd64cf211d4e

- この手順でgo.modを作成せずに`go run main.go`すると、下記のエラーが出るっぽい

`go run main.go `

```
go: cannot determine module path for source directory /home/rino/GoAndMySQL (outside GOPATH, module path must be specified)

Example usage:
        'go mod init example.com/m' to initialize a v0 or v1 module
        'go mod init example.com/m/v2' to initialize a v2 module

Run 'go help mod init' for more information.

```



`go get golang.org/x/tools/cmd/goimports`を使って
設定のformatを変えれば自動でimportしてくれる
参考
https://qiita.com/catatsuy/items/4c5a6d3efca463320b79s


- `docker-compose.yaml`ファイルについて
repositorynameはlowwer case


- コンテナ起動してcurlたたいて立ち上がってること確認

```
curl http://localhost:8080/
hello, docker container
```