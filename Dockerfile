# ---------------------------
# Dockerfile For Production
# ---------------------------

# Build Image
FROM golang:latest as build

ENV CGO_ENABLED=0 \
	GOOS=linux

WORKDIR $GOPATH/src/github.com/grandcolline/todo-list-api

COPY . .
RUN go version && go get -u -v golang.org/x/vgo
RUN vgo build


# Application Image
FROM gcr.io/distroless/base

COPY --from=build /go/src/github.com/grandcolline/todo-list-api/ /
ENV PORT=8080
EXPOSE 8080

CMD ["/todo-list-api"]

