# go-web

[![](https://img.icons8.com/color/48/000000/golang.png)](https://golang.org/)

go-web is a sample web application built using golang with a simple directory structure. 

### Features!

  - Separate ```object```, ```controller``` and ```service``` layer
  - Two-stage docker image build for go-web application
  - Docker image build and package publish using git Actions
  - Simple deployment manifest for kubernetes along with registry secret manifest
 
You can also:
  - Add repo and cache layers based on the requirement


### Packages used
 - [```gotenv```](https://pkg.go.dev/github.com/subosito/gotenv?tab=doc) - loading the environment variables to the runtime environment
 - [```gorilla/mux```](https://pkg.go.dev/github.com/gorilla/mux?tab=doc) - powerful URL router and dispatcher
 - [```rs/cors```](https://pkg.go.dev/github.com/rs/cors?tab=doc) - Cross Origin Resource Sharing W3 specification in Golang.
 - [```zap```](https://pkg.go.dev/go.uber.org/zap?tab=doc) - fast, structured, leveled logging
 - [```lumberjack.v2```](https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2?tab=doc) - rolling logger
