FROM golang:1.13.1 as builder

WORKDIR /go/src/ralali/shorty

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates


RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/ralali/shorty .

CMD ["./shorty"]