## install golang
-- download https://dl.google.com/go/go1.12.5.darwin-amd64.pkg
-- install go through pkg follow instructions

Export GOPATH
-- export PATH=$PATH:/usr/local/go/bin

validate installation here
-- go version

## install go dep
-- curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh


## To Run

1. dep init
2. dep ensure
3. go run main.go

## Run on docker

docker build -t shorty .
docker run --rm -it -p 8080:8080 shorty
