# Notes

> notes taken during the course

### 1

```sh
go version
go env
go help
```

### 3.24

```sh
go test
go mod init github.com/filipe1309/ud-go-lhc
go mod tidy
go list -m all 

go get golang.org/x/text # update dependencies

go get rsc.io/sampler
go list -m -versions rsc.io/sampler
go get rsc.io/sampler@v1.3.1 # install specific version

# Race detector
go run -race classes/sec-26-exercises-ninja-level-9-concurrency/exercise3.go              
```
