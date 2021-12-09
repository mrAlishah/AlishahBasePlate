FROM golang:1.15 as builder

ARG GITLAB_LOGIN
ARG GITLAB_TOKEN
ARG GOPRIVATE

WORKDIR /go/src/AlishahBasePlates
COPY go.mod .
COPY go.sum .

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"
ENV NO_PROXY="${GOPRIVATE}"
ENV APP_MODE="production"
ENV GOSUM="off"
ENV GIN_MODE=release
RUN echo "app mode is ${APP_MODE}"
RUN echo "machine ${GOPRIVATE} login ${GITLAB_LOGIN} password ${GITLAB_TOKEN}" > ~/.netrc
RUN go mod download


FROM builder as server_builder
WORKDIR /go/src/AlishahBasePlates

COPY . .

#RUN GIT_COMMIT=$(git rev-parse --short HEAD) \
# && BUILD_TIME=$(date +%Y/%m/%d-%H:%M:%S) \
# && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommit=$GIT_COMMIT -X main.BuildTime=$BUILD_TIME" -o server cmd/server/*.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/server/*.go


FROM debian:stretch-slim
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates
WORKDIR /usr/local/

COPY --from=server_builder /go/src/AlishahBasePlates/server .
COPY --from=server_builder /go/src/AlishahBasePlates/configs ./configs

CMD ["./server"]