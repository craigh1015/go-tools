# GO Tools

## Run and build

- go run
- go build
  - GOOS=windows go build
- go install
- go get

## Documentation

- go list
  - go list -f '{{ .Name }}: {{ .Doc }}'
  - go list -f '{{ .Imports }}'
  - go list -f '{{ .Imports }}' fmt
  - go list -f '{{ .Imports join "\n" }}' fmt
- go doc
- go doc fmt Printf
- godoc -http :6060
  - <http://localhost:6060>

## Code quality

- errcheck
  - go get -u github.com/kisielk/errcheck
  - C:\Users\craig\code\goworkspace\bin\errcheck.exe
- go vet
- go fmt
  - gofmt -d main.go
- go test

## Performance

Install Graphviz

- go-wrk
  - go get github.com/tsliwowicz/go-wrk
  - C:\Users\craig\code\goworkspace\bin\go-wrk.exe -d 5 http://localhost:8080/regexp/abc@golang.org
- _ "net/http/pprof"
- http://localhost:8080/debug/pprof/
- go tool pprof --seconds=5 localhost:8080/debug/pprof/profile
- (pprof) top
- (pprof) web
- docker run uber/go-torch -u http://10.0.0.4:8080/debug/pprof -p -t=30 > torch.svg
- go tool pprof -http=":8081" tools.exe prof.cpu

## Profile web app

``` go
package main

import (
 "fmt"
 "log"
 "net/http"

 _ "net/http/pprof"
 "regexp"
)

func main() {
 http.HandleFunc("/regexp/", handlerRegex)
 err := http.ListenAndServe(":8080", nil)
 if err != nil {
  log.Fatal(err)
 }
}
```

## Profile application

``` go
package main

import (
 "fmt"
 "log"
 "net/http"

 "regexp"

 "github.com/pkg/profile"
)

func main() {
 defer profile.Start().Stop()
 http.HandleFunc("/regexp/", handlerRegex)
 err := http.ListenAndServe(":8080", nil)
 if err != nil {
  log.Fatal(err)
 }
}
```

## grpc tools

- <https://github.com/grpc-ecosystem/awesome-grpc>
- grpcui
  - go get github.com/fullstorydev/grpcui
  - go install github.com/fullstorydev/grpcui/cmd/grpcui
  - grpcui -plaintext localhost:8080
- gzh
  - <https://ghz.sh/>
- grpc-tools
  - <https://github.com/bradleyjkemp/grpc-tools>
  - brew install bradleyjkemp/formulae/grpc-tools
  - go install github.com/bradleyjkemp/grpc-tools/...
  - grpc-dump
    - Need to use version that does not infer import paths
    - Need to use <localtest.me> rather than <localhost> to allow proxy behaviour
    - (grpc-dump) `cd grpc-tools/grpc-dump`
    - (grpc-dump) `go run main.go -port 64517 -proto_roots /Users/charris/dev/fiji-apis`
    - (hotel-search) `ENV=dev-integration http_proxy=http://localtest.me:64517 go run main.go`
    - `grpcui -plaintext localtest.me:8080`
    - `http_proxy=http://localtest.me:64517 grpcurl -d @ --plaintext localtest.me:8080 deem.hotel.searchservice.HotelService/CreateHotelSearch < CreateHotelSearch.json`
- grpc-curl
  - `grpcurl --plaintext localtest.me:8080 list`
  - `grpcurl --plaintext localtest.me:8080 describe deem.hotel.searchservice.HotelService`
