FROM golang:latest as builder

# コンテナ作業ディレクトリの変更
WORKDIR /go/src/denki/go
# ホストOSの ./go の中身を作業ディレクトリに追加
ADD ./go .

RUN go get github.com/lib/pq

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/yokoe/go-server-example
COPY . .
RUN go build main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/yokoe/go-server-example /app

CMD /app/main $PORT