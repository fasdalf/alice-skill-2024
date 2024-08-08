Training project since 2024-06-27

$ go install github.com/golang/mock/mockgen@v1.6.0
$ mockgen -source=internal/store/store.go -destination=internal/store/mock/store.go -package=mock Store
$ go mod tidy
$ go mod vendor