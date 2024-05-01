FROM golang:1.22.1 as builder
WORKDIR /src
COPY go.mod ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags 'all=-N -l' -o "example.debug" main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w' -o "example" main.go

FROM scratch as release
COPY --from=builder /src/example /usr/local/bin/example
ENTRYPOINT ["/usr/local/bin/example"]

FROM golang:1.22.1 as debug

# Install Delve
RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.1

# Copy the binary files from builder
COPY --from=builder /src/example.debug /example.debug

EXPOSE 40000
CMD ["/go/bin/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/example.debug"]
