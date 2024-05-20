# Lesson 02
Đặt tên sau

## Crypto Tools Application


* Setup Project
```shell
# init go module
go mod init myapp

# create static_variable.go file
echo "package myapp

// Version is the version of the application
var Version string
" >> static_variable.go

# create new main.go file
mkdir -p cmd/v1; echo "package main
  
import (
  \"fmt\"
    \"myapp\"
)

func main() {
    fmt.Printf(\"Crypto Tools Application\n------------------------------------\n\"+
    \"\tVersion: %s\n------------------------------------\n\", myapp.Version)
}
" >> cmd/v1/main.go
```
* Build app
```shell
go mod tidy;\
  go build -ldflags "-s -w -extldflags '-static' -X myapp.Version=beta-1.0.0" \
  -o bin/myapp \
  -trimpath cmd/v1/main.go
```
* Run build
```shell
./bin/myapp
```

