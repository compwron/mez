# from https://blog.golang.org/docker

# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/golang/compwron/mez

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/compwron/mez

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/outyet

# Document the listening port
EXPOSE 3000