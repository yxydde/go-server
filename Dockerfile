# Build
FROM golang:1.15.2-alpine as build

RUN apk add --no-cache git 

RUN mkdir -p /go/src/github.com/yxydde

WORKDIR /go/src/github.com/yxydde

RUN git clone https://github.com/yxydde/go-server.git

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN cd go-server && go build


# Run
FROM alpine:3.12 as prod

RUN mkdir -p /app/conf

WORKDIR /app

COPY --from=build /go/src/github.com/yxydde/go-server/go-server /app

COPY --from=build /go/src/github.com/yxydde/go-server/conf /app/conf

ENV Version=v2.0

expose 8080

CMD ["/app/go-server"]