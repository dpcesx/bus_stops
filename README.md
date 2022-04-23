# README

All commands start from root `.`

```sh

go version
# go version go1.18.1 darwin/amd64

node --version
# v16.13.1

cd bus_stop
go test

cd service
go mod tidy
go get .
go run .

curl localhost:8080/stop/1

cd client
npm start
open http://localhost:8000/

```