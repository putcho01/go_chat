FROM golang:1.17.7-alpine
ENV TZ=Asia/Tokyo

# ENV CGO_ENABLED=0
# ENV GOOS=linux
# ENV GOARCH=amd64
# ENV GOPATH /go
# ENV GO111MODULE on

RUN apk update && \
  apk --no-cache add git

RUN mkdir /go/app
WORKDIR /go/app
COPY . /go/app

# 必要なパッケージをインストール
RUN go mod tidy
# Airをインストール
RUN go install github.com/cosmtrek/air@v1.40.0

EXPOSE 80
CMD ["air", "-c", ".air.toml"]