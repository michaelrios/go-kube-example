FROM golang:1.10.1-alpine3.7
RUN apk add --no-cache git
RUN apk add --update make
RUN go get github.com/golang/dep/cmd/dep
RUN go get -u golang.org/x/lint/golint

# define work directory
RUN mkdir -p /go/src/github.com/michaelrios/go-kube-example/
WORKDIR /go/src/github.com/michaelrios/go-kube-example/

# Get Dependencies
COPY Gopkg.lock Gopkg.toml /go/src/github.com/michaelrios/go-kube-example/
RUN dep ensure -vendor-only

# Adding source files
COPY . /go/src/github.com/michaelrios/go-kube-example/

# Ensure code quality
RUN make lint
RUN make test

# serve the app
RUN go build main.go
EXPOSE 80
CMD ["./main"]
