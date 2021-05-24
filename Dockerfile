FROM golang:1.14-alpine
WORKDIR /usr/src/go-skeleton
ENV TZ Asia/Shanghai
ENV GOPROXY=https://goproxy.cn
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN go build main.go
CMD ["/usr/src/go-skeleton/main"]
