ARG GO_VERSION=latest

# ------------------------
# Develop Stage
# ------------------------
FROM golang:${GO_VERSION} as dev

ARG REALIZE_VERSION=v2.0.2
ARG GOLANGCI_VERSION=v1.18.0

ENV GOOS=linux
ENV GO111MODULE=on

# install development tools
WORKDIR $GOPATH/src/tools
RUN go get \
	gopkg.in/urfave/cli.v2@master \
	github.com/oxequa/realize@${REALIZE_VERSION} \
	golang.org/x/tools/cmd/goimports@latest \
	github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_VERSION}
RUN rm -rf /go/pkg && rm -rf /go/src/*

WORKDIR $GOPATH/src/github.com/grandcolline/todo-list-api

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN env CGO_ENABLED=0 go install


# ------------------------
# Runtime Stage
# ------------------------
FROM gcr.io/distroless/static as run

COPY --from=dev /go/bin/todo-list-api /todo-list-api
CMD ["/todo-list-api"]
